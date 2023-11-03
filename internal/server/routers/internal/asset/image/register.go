package image

import "github.com/gin-gonic/gin"

func RegisterImageApi(group *gin.RouterGroup) {
	group.POST("/image/upload", UploadImage)
	group.GET("/image/fetch", FetchImage)
}
