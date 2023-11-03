package models

import (
	"gorm.io/gorm"
	"time"
)

const ProductTable string = "product"

type Product struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	ImageId       string  `gorm:"not null" json:"image_id" binding:"required"`
	User          string  `gorm:"not null" json:"user"`
	Name          string  `gorm:"not null" json:"name" binding:"required"`
	Summary       string  `gorm:"not null" json:"summary" binding:"required"`
	Price         float64 `gorm:"not null" json:"price" binding:"required"`
	DiscountPrice float64 `gorm:"default:0" json:"discount_price"`

	Published  int `gorm:"default:0" json:"published"`
	CanPublish int `gorm:"default:1" json:"can_publish"`
	Sold       int `gorm:"default:0" json:"sold"`
}
