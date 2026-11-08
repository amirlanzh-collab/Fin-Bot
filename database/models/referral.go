package models

import (
	"time"
)

type Referral struct {
	ID         uint `gorm:"primaryKey"`
	ReferrerID uint `gorm:"not null;index;check:referrer_id <> invited_id"`
	InvitedID  uint `gorm:"not null;uniqueIndex"`
	CreatedAt  time.Time
}
