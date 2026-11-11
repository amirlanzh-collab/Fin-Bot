package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Referral struct {
	ID         uint `gorm:"primaryKey"`
	ReferrerID uint `gorm:"not null;index;check:referrer_id <> invited_id"`
	InvitedID  uint `gorm:"not null;uniqueIndex"`
	CreatedAt  time.Time
}

type ReferralService struct {
	db *gorm.DB
}

func NewReferralService(db *gorm.DB) *ReferralService {
	return &ReferralService{db: db}
}

func createRefBonus(accountID uint, db *gorm.DB, amount decimal.Decimal) error {
	operation := &Operation{
		AccountID: accountID,
		Amount:    amount,
		Type:      OpReferralBonus,
	}
	err := db.Create(operation).Error
	if err != nil {
		return err
	}

	err = db.Model(&Account{}).
		Where("id = ?", accountID).
		UpdateColumn("balance", gorm.Expr("balance + ?", operation.Amount)).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ReferralService) GetRefList(tgID int64) ([]User, error) {
	var user User

	err := s.db.Where("tg_id = ?", tgID).First(&user).Error
	if err != nil {
		return nil, err
	}

	var invited []User

	err = s.db.
		Joins("JOIN referrals ON referrals.invited_id = users.id").
		Where("referrals.referrer_id = ?", user.ID).
		Order("referrals.created_at DESC").
		Find(&invited).Error
	if err != nil {
		return nil, err
	}
	return invited, nil
}
