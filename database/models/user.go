package models

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID         uint      `gorm:"primaryKey"`
	TgID       int64     `gorm:"uniqueIndex;not null"`
	Username   string    `gorm:"size(32);index;default:''"`
	ReferrerID *uint     `gorm:"index"`
	Accounts   []Account `gorm:"foreignKey:UserID;constraint:OnDelete:RESTRICT"`
	CreatedAt  time.Time
}

type UserService struct {
	db            *gorm.DB
	welcomeBonus  decimal.Decimal
	referralBonus decimal.Decimal
}

func NewUserService(db *gorm.DB, welcomeBonus decimal.Decimal, referralBonus decimal.Decimal) *UserService {
	return &UserService{db: db, welcomeBonus: welcomeBonus, referralBonus: referralBonus}
}

func (s *UserService) RegisterUser(tgID int64, username string, referrerTgID int64) (*User, error) {
	var user *User

	err := s.db.Transaction(func(trs *gorm.DB) error {
		var err error
		user, created, err := createUser(tgID, username, trs)
		if err != nil {
			return err
		}

		if !created {
			return nil
		}

		if referrerTgID != 0 && referrerTgID != tgID {
			var referrer User

			err := trs.Where("tg_id = ?", referrerTgID).First(&referrer).Error
			if err == nil {
				err = trs.Model(user).Update("referrer_id", referrer.ID).Error
				if err != nil {
					return err
				}
				referral := &Referral{ReferrerID: referrer.ID, InvitedID: user.ID}
				err = trs.Create(referral).Error
				if err != nil {
					return err
				}

				err = createRefBonus(referrer.ID, trs, s.referralBonus)
				if err != nil {
					return err
				}
			} else if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
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

func createUser(tgID int64, username string, db *gorm.DB) (*User, bool, error) {
	user := &User{TgID: tgID, Username: username}

	res := db.Clauses(clause.OnConflict{DoNothing: true}).Create(user)
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
