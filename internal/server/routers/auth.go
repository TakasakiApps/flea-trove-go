package routers

import (
	"github.com/TakasakiApps/flea-trove-go/internal/database"
	"github.com/TakasakiApps/flea-trove-go/internal/models"
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/hanakogo/digine"
	"github.com/hanakogo/exceptiongo"
	"github.com/hanakogo/hanakoutilgo"
	"net/http"
)

func registerAuthApi() {
	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/register", register)
	authGroup.POST("/login", login)
}

var register gin.HandlerFunc = func(c *gin.Context) {
	// 解析传过来的对象
	user := utils.CtxBindJson[models.User](c)
	// 查找用户是否已经被注册过
	if query := database.User().GetUserByAccount(user.Account); query != nil {
		exceptiongo.ThrowMsg[types.StatusConflict]("用户已存在")
	}
	// 加密密码
	user.Password = utils.BCryptPassword(user.Password)
	// 新增用户
	effected := database.User().AddUser(user)

	// 检查受影响行数是否是1，如果是其他数字，则SQL没有执行成功
	if effected != 1 {
		exceptiongo.ThrowMsg[types.StatusInternalServerError]("注册失败")
	}

	utils.CtxRespOK[string](c, "注册成功")
}

var login gin.HandlerFunc = func(c *gin.Context) {
	// 解析传过来的对象
	userLogin := utils.CtxBindJson[models.UserLogin](c)
	// 拿到JWT Secret
	jwtSecret := *digine.Require[string](digine.NewLabel("JWT_SECRET"))

	// 封装一个生成7天有效期的token并返回前端的函数
	genToken := func(user models.User) {
		token := utils.JWTSign[models.User](user, jwtSecret, 60*60*24*7)
		utils.CtxRespOK[map[string]string](c, map[string]string{
			"token": token,
		})
	}

	// 如果请求token鉴权
	if userLogin.TokenAuth {
		// 从上下文中获取token
		token, exists := c.Get("token")
		if !exists {
			exceptiongo.ThrowMsg[types.StatusUnauthorized]("Token不存在")
		}
		// 校验token
		utils.JWTParse[models.User](
			hanakoutilgo.CastTo[string](token), jwtSecret,
			func(isValid bool, data models.User) {
				if !isValid {
					exceptiongo.ThrowMsg[types.StatusUnauthorized]("Token无效")
				}
				userQuery := database.User().GetUserByAccount(data.Account)
				// 生成新token返回（续期）
				genToken(*userQuery)
			},
		)
		// 停止执行写在外部是因为，这个判断中已经是参数中确定是以token方式鉴权了
		// 所以不需要再去用户名密码鉴权
		c.Abort()
		return
	}

	// 如果是用户名密码方式鉴权
	// 先查询用户是否存在
	if query := database.User().GetUserByAccount(userLogin.Account); query != nil {
		// 利用工具函数，比对bcrypt加密过的密码
		if utils.CompareP2B(userLogin.Password, query.Password) {
			// 生成token并返回
			genToken(*query)
			c.Abort()
			return
		}
	} else {
		exceptiongo.ThrowMsg[types.StatusUnauthorized]("用户不存在")
	}

	// 如果没有鉴权通过，则返回密码错误
	utils.CtxRespCustom[string](c, http.StatusUnauthorized, "密码错误")
}
