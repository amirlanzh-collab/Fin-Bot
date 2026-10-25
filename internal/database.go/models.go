package database

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

type User struct {
	ID         uint      `gorm:"primaryKey"`
	TgID       int64     `gorm:"uniqueIndex;not null"`
	ReferrerID *uint     `gorm:"index"`
	Referral   *User     `gorm:"foreignKey:ReferrerID;constraint:OnDelete:RESTRICT"`
	Accounts   []Account `gorm:"foreignKey:UserID;constraint:OnDelete:RESTRICT"`
	CreatedAt  time.Time
}

type Account struct {
	ID         uint            `gorm:"primaryKey"`
	UserID     uint            `gorm:"not null;index"`
	Balance    decimal.Decimal `gorm:"type:numeric(20,2);not null;default:0;check:balance >= 0"`
	Operations []Operation     `gorm:"foreignKey:AccountID;constraint:OnDelete:RESTRICT"`
}

type Operation struct {
	ID        uint            `gorm:"primaryKey"`
	AccountID uint            `gorm:"not null;index:idx_op_account_created,priority:1"`
	Amount    decimal.Decimal `gorm:"type:numeric(20,2);not null"`
	Type      OperationType   `gorm:"type:varchar(32);not null;check:type IN ('welcome_bonus','referral_bonus','transfer')"`
	CreatedAt time.Time       `gorm:"not null;index:idx_op_account_created,priority:2,sort:desc"`
}

type Referral struct {
	ID         uint `gorm:"primaryKey"`
	ReferrerID uint `gorm:"not null;index;check:referrer_id <> invited_id"`
	InvitedID  uint `gorm:"not null;uniqueIndex"`
	CreatedAt  time.Time
}
