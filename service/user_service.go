package service

import (
	"errors"

	"finbot/database/models"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) RegisterUser(tgID int64) (*models.User, error) {
	var user *models.User

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

		if err := createWelcomeBonus(account.ID, trs); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func createUser(tgID int64, db *gorm.DB) (*models.User, error) {
	user := &models.User{TgID: tgID}
	user.TgID = tgID
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func createAccount(userID uint, db *gorm.DB) (*models.Account, error) {
	account := &models.Account{UserID: userID}
	err := db.Create(&account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func createWelcomeBonus(accountID uint, db *gorm.DB) error {
	operation := models.Operation{
		AccountID: accountID,
		Amount:    decimal.NewFromInt(1000),
		Type:      models.OpWelcomeBonus,
	}
	err := db.Create(&operation).Error
	if err != nil {
		return err
	}

	err = db.Model(&models.Account{}).
		Where("id = ?", accountID).
		UpdateColumn("balance", gorm.Expr("balance + ?", operation.Amount)).Error
	if err != nil {

		return err
	}
	return nil
}
