package models

import (
	"gorm.io/gorm"
	"time"
)

const OrderTable string = "order"

type Order struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	User      string `gorm:"not null" json:"user" binding:"required"`
	ProductId uint   `gorm:"not null" json:"product_id" binding:"required"`

	CustomerName string `gorm:"not null" json:"customer_name" binding:"required"`
	Address      string `gorm:"not null" json:"address" binding:"required"`
	Phone        string `gorm:"not null" json:"phone" binding:"required"`
	Remark       string `json:"remark"`

	Paid   int `gorm:"default:0" json:"paid"`
	Closed int `gorm:"default:0" json:"closed"`
}
