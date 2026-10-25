package service

import (
	"errors"
	"finbot/internal/database.go"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) RegisterUser(tgID int64) (*database.User, error) {
	var user database.User

	err := s.db.Where("tg_id = ?", tgID).First(&user).Error
	if err == nil {
		return &user, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = s.db.Transaction(func(trs *gorm.DB) error {
		user.TgID = tgID
		err = trs.Create(&user).Error
		if err != nil {
			return err
		}

		account := database.Account{UserID: user.ID}
		err = trs.Create(&account).Error
		if err != nil {
			return err
		}

		operation := database.Operation{
			AccountID: account.ID,
			Amount:    decimal.NewFromInt(1000),
			Type:      database.OpWelcomeBonus,
		}
		err = trs.Create(&operation).Error
		if err != nil {
			return err
		}

		err = trs.Model(&database.Account{}).
			Where("id = ?", account.ID).
			UpdateColumn("balance", gorm.Expr("balance + ?", operation.Amount)).Error
		if err != nil {

			return err
		}
		return nil
	})

	return &user, nil
}
