package models

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	ID         uint      `gorm:"primaryKey"`
	TgID       int64     `gorm:"uniqueIndex;not null"`
	ReferrerID *uint     `gorm:"index"`
	Referral   *User     `gorm:"foreignKey:ReferrerID;constraint:OnDelete:RESTRICT"`
	Accounts   []Account `gorm:"foreignKey:UserID;constraint:OnDelete:RESTRICT"`
	CreatedAt  time.Time
}

type UserService struct {
	db           *gorm.DB
	welcomeBonus decimal.Decimal
}

func NewUserService(db *gorm.DB, welcomeBonus decimal.Decimal) *UserService {
	return &UserService{db: db, welcomeBonus: welcomeBonus}
}

func (s *UserService) RegisterUser(tgID int64) (*User, error) {
	var user *User

	err := s.db.Where("tg_id = ?", tgID).First(&user).Error
	if err == nil {
		return user, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = s.db.Transaction(func(trs *gorm.DB) error {
		var err error
		user, err = createUser(tgID, trs)
		if err != nil {
			return err
		}

		account, err := createAccount(user.ID, trs)
		if err != nil {
			return err
		}

		if err := createWelcomeBonus(account.ID, trs, s.welcomeBonus); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func createUser(tgID int64, db *gorm.DB) (*User, error) {
	user := &User{TgID: tgID}
	user.TgID = tgID
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func createAccount(userID uint, db *gorm.DB) (*Account, error) {
	account := &Account{UserID: userID}
	err := db.Create(&account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func createWelcomeBonus(accountID uint, db *gorm.DB, amount decimal.Decimal) error {
	operation := Operation{
		AccountID: accountID,
		Amount:    amount,
		Type:      OpWelcomeBonus,
	}
	err := db.Create(&operation).Error
	if err != nil {
		return err
	}

	err = db.Model(Account{}).
		Where("id = ?", accountID).
		UpdateColumn("balance", gorm.Expr("balance + ?", operation.Amount)).Error
	if err != nil {

		return err
	}
	return nil
}
