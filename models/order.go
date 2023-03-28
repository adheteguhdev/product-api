package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	ID        uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Name      string         `gorm:"size:255;unique_index;not null" json:"name"`
	Price     float64        `gorm:"not null;type:decimal(10,2)" json:"price"`
	ExpiredAt time.Time      `gorm:"default:null" json:"expired_at"`
	CreatedAt time.Time      `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type OrderHistory struct {
	ID           uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserID       uint64    `gorm:"not null" json:"user_id"`
	OrderItemID  uint64    `gorm:"not null" json:"order_item_id"`
	Descriptions string    `gorm:"type:text" json:"descriptions"`
	CreatedAt    time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp" json:"updated_at"`

	User      User      `gorm:"foreignkey:UserID"`
	OrderItem OrderItem `gorm:"foreignkey:OrderItemID"`
}
