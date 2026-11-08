package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Account struct {
	ID         uint            `gorm:"primaryKey"`
	UserID     uint            `gorm:"not null;index"`
	Balance    decimal.Decimal `gorm:"type:numeric(20,2);not null;default:0;check:balance >= 0"`
	Operations []Operation     `gorm:"foreignKey:AccountID;constraint:OnDelete:RESTRICT"`
}

type AccountService struct {
	db *gorm.DB
}

func NewAccounrService(db *gorm.DB) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) GetBalance(tgID int64) (decimal.Decimal, error) {
	var account Account
	err := s.db.
		Joins("JOIN users ON users.id = accounts.user_id").
		Where("users.tg_id = ?", tgID).
		First(&account).Error

	if err != nil {
		return decimal.Zero, err
	}

	return account.Balance, nil
}
