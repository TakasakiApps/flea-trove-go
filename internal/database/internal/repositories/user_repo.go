package repositories

import (
	"github.com/TakasakiApps/flea-trove-go/internal/models"
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/hanakogo/exceptiongo"
)

type UserRepository BaseRepository

// GetUserById 根据ID获取用户
func (ur *UserRepository) GetUserById(id int) (user *models.User) {
	var userRes models.User
	tx := ur.db.Table(models.UserTable).Where("id = ?", id).First(&userRes)
	if tx.RowsAffected != 1 {
		return nil
	}
	user = &userRes
	return
}

// GetUserByAccount 根据账号获取用户
func (ur *UserRepository) GetUserByAccount(account string) *models.User {
	var users []*models.User
	ur.db.Table(models.UserTable).Where("account = ?", account).First(&users)

	if len(users) == 0 {
		return nil
	}
	if len(users) > 1 {
		exceptiongo.ThrowMsgF[types.DBDuplicateDataError]("too many record in table %v", models.UserTable)
	}

	return users[0]
}

// AddUser 增加用户
func (ur *UserRepository) AddUser(user *models.User) int64 {
	result := ur.db.Table(models.UserTable).Create(user)

	return result.RowsAffected
}
