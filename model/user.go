package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Active bool   `json:"active"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
