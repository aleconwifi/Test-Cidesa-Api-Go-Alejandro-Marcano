package model

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID uint `gorm:"primarykey"`

	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
	Orders    []Order `json:"orders" gorm:"many2many:items_orders;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ItemsOrders struct {
	ID uint `gorm:"primarykey"`

	ItemID uint64 `json:"itemID"`
	Item   Item   `json:"item" gorm:"foreignKey:ItemID"`

	OrderID uint64 `json:"orderID"`
	Order   Order  `json:"order" gorm:"foreignKey:OrderID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
