package repositories

import (
	"github.com/hanakogo/digine"
	"gorm.io/gorm"
)

func BindRepositories(db *gorm.DB) {
	digine.Bind[UserRepository](&UserRepository{db}, digine.NilLabel)
	digine.Bind[ProductRepository](&ProductRepository{db}, digine.NilLabel)
	digine.Bind[OrderRepository](&OrderRepository{db}, digine.NilLabel)
}
