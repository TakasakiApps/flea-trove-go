package initial

import (
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gookit/goutil/fsutil"
	"github.com/hanakogo/exceptiongo"
	"os"
)

func ensureAppDataHome() {
	appDataHomePath := utils.GetAppDataHome()
	// 如果数据目录是一个文件，则抛出异常
	if fsutil.FileExists(appDataHomePath) {
		exceptiongo.ThrowMsgF[types.Fatal]("%v seems is a file, please remove it then try again", appDataHomePath)
	}
	// 如果数据目录不存在，则创建
	if !fsutil.DirExist(appDataHomePath) {
		err := fsutil.Mkdir(appDataHomePath, os.ModeDir)
		exceptiongo.ThrowErr[types.IOError](err)
	}
}
