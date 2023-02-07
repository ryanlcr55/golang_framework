package gormRepo

import (
	"gorm.io/gorm"
)

type TransactionHandler struct {
	DB *gorm.DB
}

func (h TransactionHandler) Begin() any {
	h.DB = h.DB.Begin()
	return h.DB
}

func (h TransactionHandler) RollBack() {
	h.DB.Rollback()
}

func (h TransactionHandler) Commit() error {
	return h.DB.Commit().Error
}
