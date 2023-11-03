package models

import (
	"gorm.io/gorm"
	"time"
)

const UserTable string = "user"

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Username string `json:"username"`
	Account  string `gorm:"not null" json:"account"`
	Password string `gorm:"not null" json:"password"`
}

type UserLogin struct {
	User
	TokenAuth bool `json:"token_auth"`
}
