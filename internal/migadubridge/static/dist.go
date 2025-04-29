//go:build !dev

//生产模式下的嵌入文件

package static

import (
	"embed"

	"github.com/gin-contrib/static"
)

//go:embed dist
var server embed.FS

func GetFS() (static.ServeFileSystem, error) {
	return static.EmbedFolder(server, "dist")
}
