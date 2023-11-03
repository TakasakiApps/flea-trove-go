package utils

import (
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gookit/goutil/fsutil"
	"github.com/hanakogo/exceptiongo"
	"os"
	"path/filepath"
)

func ensureImageAssetsDir() string {
	path := GetAppDataFile("upload", "images")
	if !fsutil.DirExist(path) {
		err := fsutil.Mkdir(path, os.ModeDir)
		exceptiongo.ThrowErr[types.IOError](err)
	}
	return path
}

func getImagePath(id string) string {
	baseImageDir := ensureImageAssetsDir()
	imagePath := filepath.Join(baseImageDir, id)
	return imagePath
}

func ImageExists(id string) bool {
	return fsutil.FileExists(getImagePath(id))
}

func ImageSave(c *gin.Context) (string, error) {
	_, headers, _ := c.Request.FormFile("file")
	id := uuid.New().String()
	imagePath := getImagePath(id)
	return id, c.SaveUploadedFile(headers, imagePath)
}

func ImagePath(id string) string {
	return getImagePath(id)
}
