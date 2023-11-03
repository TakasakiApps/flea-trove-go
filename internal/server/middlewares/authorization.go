package middlewares

import (
	"github.com/TakasakiApps/flea-trove-go/internal/database"
	"github.com/TakasakiApps/flea-trove-go/internal/models"
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/arrutil"
	"github.com/hanakogo/digine"
	"github.com/hanakogo/exceptiongo"
	"strings"
)

func TokenExtractor(c *gin.Context) string {
	// 如果路径里有token，直接拿出来
	token := c.Query("token")
	if token != "" {
		return token
	}

	// 如果路径里没有token，从header里提取
	authorization := c.Request.Header.Get("Authorization")
	if len(strings.Split(authorization, " ")) == 2 {
		return strings.Split(authorization, " ")[1]
	}
	return ""
}

func getAuthorization(skipApi ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 提取请求传过来的token
		token := TokenExtractor(c)
		// 把token存到上下文中
		c.Set("token", token)

		// 如果API是无需鉴权的
		if arrutil.HasValue(skipApi, c.FullPath()) {
			c.Next()
			return
		}

		// 如果需要鉴权，开始校验token
		utils.JWTParse[models.User](
			token, *digine.Require[string](digine.NewLabel("JWT_SECRET")),
			func(isValid bool, data models.User) {
				if !isValid {
					exceptiongo.ThrowMsg[types.StatusUnauthorized]("未登录，无权限调用")
				}
				if database.User().GetUserByAccount(data.Account) == nil {
					exceptiongo.ThrowMsg[types.StatusUnauthorized]("用户不存在，无法校验token")
				}
				c.Set("user_account", data.Account)
			},
		)
	}
}
