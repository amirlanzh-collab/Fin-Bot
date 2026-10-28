package service

import (
	"finbot/internal/database.go"

	"gorm.io/gorm"
)

type OperationService struct {
	db *gorm.DB
}

func NewOperationService(db *gorm.DB) *OperationService {
	return &OperationService{db: db}
}

func (s *OperationService) GetHistory(tgID int64) ([]database.Operation, error) {
	var operations []database.Operation

	err := s.db.
		Joins("JOIN accounts ON accounts.id =  operations.account_id").
		Joins("JOIN users ON users.id = accounts.user_ID").
		Order("operations.created_at DESC, operations.id DESC").
		Where("users.tg_id = ?", tgID).
		Find(&operations).Error

	if err != nil {
		return nil, err
	}
	return operations, nil
}
