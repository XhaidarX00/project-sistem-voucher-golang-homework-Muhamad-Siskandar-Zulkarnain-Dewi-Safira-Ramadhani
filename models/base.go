package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}