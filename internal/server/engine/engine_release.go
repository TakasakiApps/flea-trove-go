//go:build release

package engine

import "github.com/gin-gonic/gin"

func setMode() {
	gin.SetMode("release")
}
