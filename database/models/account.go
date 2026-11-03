package models

import (
	"github.com/shopspring/decimal"
)

type Account struct {
	ID         uint            `gorm:"primaryKey"`
	UserID     uint            `gorm:"not null;index"`
	Balance    decimal.Decimal `gorm:"type:numeric(20,2);not null;default:0;check:balance >= 0"`
	Operations []Operation     `gorm:"foreignKey:AccountID;constraint:OnDelete:RESTRICT"`
}
