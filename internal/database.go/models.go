package database

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TgID       int64       `gorm:"uniqueIndex;not null"`
	InvitedBy  *int64      `gorm:"index"` //тг ID того кто пригласил
	Ballancess []Ballance  `gorm:"foreignKey:UserID"`
	Operations []Operation `gorm:"foreignKey:UserID"`
}

type Ballance struct {
	gorm.Model
	UserID   uint            `gorm:"not null;index"`
	Ballance decimal.Decimal `gorm:"type:numeric(15,2);default:0.00;not null"`
}

type Operation struct {
	gorm.Model
	UserID    uint            `gorm:"not null;index"`
	Amount    decimal.Decimal `gorm:"type:numeric(15,2);not null"`
	Type      string          `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
}
