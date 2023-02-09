package gormRepo

import (
	"gorm.io/gorm"
)

type TrxHandler struct {
	DB *gorm.DB
}

func (h TrxHandler) Begin() any {
	h.DB = h.DB.Begin()
	return h.DB
}

func (h TrxHandler) RollBack() {
	h.DB.Rollback()
}

func (h TrxHandler) Commit() error {
	return h.DB.Commit().Error
}

func NewTransactionHandler(db *gorm.DB) *TrxHandler {
	return &TrxHandler{}
}
