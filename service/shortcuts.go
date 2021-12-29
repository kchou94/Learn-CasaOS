package service

import "gorm.io/gorm"

type ShortcutsService interface{}

type shortcutsService struct {
	db *gorm.DB
}

func NewShortcutsService(db *gorm.DB) ShortcutsService {
	return &shortcutsService{db: db}
}
