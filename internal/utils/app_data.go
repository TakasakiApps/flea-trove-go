package utils

import (
	"github.com/TakasakiApps/flea-trove-go/internal/consts"
	"github.com/gookit/goutil/sysutil"
	"path/filepath"
)

func GetAppDataHome() string {
	return filepath.Join(sysutil.HomeDir(), "."+consts.AppName)
}

func GetAppDataFile(paths ...string) (path string) {
	path = GetAppDataHome()
	for _, p := range paths {
		path = filepath.Join(path, p)
	}
	return
}
