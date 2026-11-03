package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type OperationType string

const (
	OpWelcomeBonus  OperationType = "welcome_bonus"
	OpRefferalBonus OperationType = "refferal_bonus"
	OpTransfer      OperationType = "transfer"
)

type Operation struct {
	ID        uint            `gorm:"primaryKey"`
	AccountID uint            `gorm:"not null;index:idx_op_account_created,priority:1"`
	Amount    decimal.Decimal `gorm:"type:numeric(20,2);not null"`
	Type      OperationType   `gorm:"type:varchar(32);not null;check:type IN ('welcome_bonus','referral_bonus','transfer')"`
	CreatedAt time.Time       `gorm:"not null;index:idx_op_account_created,priority:2,sort:desc"`
}
