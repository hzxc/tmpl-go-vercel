package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// tag json:"-" skip this filed
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
