package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID            uint    `gorm:"primarykey"`
	NROOrder      string  `json:"NROOrder"`
	UserID        uint64  `json:"userID"`
	User          User    `json:"user" gorm:"foreignKey:UserID"`
	PromotionID   uint64  `json:"promotionID"`
	Status        string  `json:"status"`
	Total         float64 `json:"total"`
	TotalDiscount float64 `json:"totalDiscount"`

	Items []Item `json:"items" gorm:"many2many:items_orders;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
