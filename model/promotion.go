package model

import (
	"time"

	"gorm.io/gorm"
)

type Promotion struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name"`
	Code string `json:"code"`
	Used bool   `json:"used"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
