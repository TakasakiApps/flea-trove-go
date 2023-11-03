package database

import (
	"github.com/TakasakiApps/flea-trove-go/internal/database/internal/repositories"
	"github.com/hanakogo/digine"
)

func User() *repositories.UserRepository {
	return digine.Require[repositories.UserRepository](digine.NilLabel)
}

func Product() *repositories.ProductRepository {
	return digine.Require[repositories.ProductRepository](digine.NilLabel)
}

func Order() *repositories.OrderRepository {
	return digine.Require[repositories.OrderRepository](digine.NilLabel)
}
