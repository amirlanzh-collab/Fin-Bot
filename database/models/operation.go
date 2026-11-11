package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OperationType string

const (
	OpWelcomeBonus  OperationType = "welcome_bonus"
	OpReferralBonus OperationType = "referral_bonus"
	OpTransfer      OperationType = "transfer"
)

type Operation struct {
	ID        uint            `gorm:"primaryKey"`
	AccountID uint            `gorm:"not null;index:idx_op_account_created,priority:1"`
	Amount    decimal.Decimal `gorm:"type:numeric(20,2);not null"`
	Type      OperationType   `gorm:"type:varchar(32);not null;check:type IN ('welcome_bonus','referral_bonus','transfer')"`
	CreatedAt time.Time       `gorm:"not null;index:idx_op_account_created,priority:2,sort:desc"`
}

type OperationService struct {
	db *gorm.DB
}

func NewOperationService(db *gorm.DB) *OperationService {
	return &OperationService{db: db}
}

func (s *OperationService) GetHistory(tgID int64) ([]Operation, error) {
	var operations []Operation

	err := s.db.
		Joins("JOIN accounts ON accounts.id = operations.account_id").
		Joins("JOIN users ON users.id = accounts.user_ID").
		Order("operations.created_at DESC, operations.id DESC").
		Where("users.tg_id = ?", tgID).
		Find(&operations).Error

	if err != nil {
		return nil, err
	}
	return operations, nil
}
