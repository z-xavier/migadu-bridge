//go:build dev

// 开发模式下的文件系统读取

package static

import (
	"github.com/gin-contrib/static"
)

// 定义前端静态文件目录路径
const frontendDistPath = "./frontend/dist"

func GetFS() (static.ServeFileSystem, error) {
	return static.LocalFile(frontendDistPath, false), nil
}
