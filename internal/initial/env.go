package initial

import (
	"github.com/gookit/goutil/strutil"
	"github.com/hanakogo/digine"
	"os"
)

func initEnv() {
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		jwtKey = "default_secret"
	}
	digine.Bind[string](&jwtKey, digine.NewLabel("JWT_SECRET"))

	portStr := os.Getenv("SERVER_PORT")
	port := strutil.IntOr(portStr, 8080)
	digine.Bind[int](&port, digine.NewLabel("SERVER_PORT"))
}
