package engine

import "github.com/gin-gonic/gin"

var Gin *gin.Engine

func init() {
	setMode()
	Gin = gin.Default()
}
