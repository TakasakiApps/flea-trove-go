package database

import (
	"github.com/TakasakiApps/flea-trove-go/internal/database/internal/repositories"
	"github.com/TakasakiApps/flea-trove-go/internal/models"
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/gookit/slog"
	"github.com/hanakogo/digine"
	"github.com/hanakogo/exceptiongo"
	"gorm.io/gorm"
)

var logger *slog.Logger
var db *gorm.DB

func Migrate() {
	logger = digine.Require[slog.Logger](digine.NilLabel)
	db = digine.Require[gorm.DB](digine.NilLabel)

	// 数据库迁移
	doMigrating()
	// 初始化DAO
	repositories.BindRepositories(db)
}

func doMigrating() {
	logger.Info("Start migrating database...")

	migrator := digine.Require[gorm.DB](digine.NilLabel).Migrator()

	targets := []types.Pair[any, string]{
		{&models.User{}, models.UserTable},
		{&models.Product{}, models.ProductTable},
		{&models.Order{}, models.OrderTable},
	}

	for _, obj := range targets {
		if !migrator.HasTable(obj.First) && !migrator.HasTable(obj.Second) {
			logger.Infof("Migrating dataBase<%s>", obj.Second)
			err := migrator.CreateTable(obj.First)
			exceptiongo.ThrowErr[types.DBIOError](err)
			if !migrator.HasTable(obj.Second) {
				err = migrator.RenameTable(obj.First, obj.Second)
				exceptiongo.ThrowErr[types.DBIOError](err)
			}
		}
	}
}
