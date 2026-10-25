package service

import (
	"finbot/internal/database.go"

	"github.com/shopspring/decimal"

	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
}

func NewAccounrService(db *gorm.DB) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) GetBalance(tgID int64) (decimal.Decimal, error) {
	var account database.Account
	err := s.db.
		Joins("JOIN users ON users.id = accounts.user_id").
		Where("users.tg_id = ?", tgID).
		First(&account).Error

	if err != nil {
		return decimal.Zero, err
	}

	return account.Balance, nil
}
