package models

import (
	"time"
)

type User struct {
	ID         uint      `gorm:"primaryKey"`
	TgID       int64     `gorm:"uniqueIndex;not null"`
	ReferrerID *uint     `gorm:"index"`
	Referral   *User     `gorm:"foreignKey:ReferrerID;constraint:OnDelete:RESTRICT"`
	Accounts   []Account `gorm:"foreignKey:UserID;constraint:OnDelete:RESTRICT"`
	CreatedAt  time.Time
}
