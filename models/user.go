package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint64         `gorm:"primary_key:auto_increment" json:"id"`
	FullName   string         `gorm:"size:255;not null" json:"full_name"`
	FirstOrder *uint64         `gorm:"default:null" json:"first_order"`
	CreatedAt  time.Time      `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"default:current_timestamp" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`

	// OrderHistory []OrderHistory `gorm:"foreignkey:UserID"`
}
