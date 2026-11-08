package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	err := s.db.Transaction(func(trs *gorm.DB) error {
		var err error
		user, created, err := createUser(tgID, trs)
		if err != nil {
			return err
		}

		if !created {
			return nil
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

func createUser(tgID int64, db *gorm.DB) (*User, bool, error) {
	user := &User{TgID: tgID}

	res := db.Clauses(clause.OnConflict{DoNothing: true}).Create(user).First(user)
	if res.Error != nil {
		return nil, false, res.Error
	}
	if res.RowsAffected == 0 {
		err := db.Where("tg_id = ?", tgID).First(user).Error
		if err != nil {
			return nil, false, err
		}
		return user, false, nil
	}

	return user, true, nil
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
