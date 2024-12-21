package migadubridge

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/core"
	"migadu-bridge/internal/errmsg"
	"migadu-bridge/internal/log"
)

func installInteriorWebRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errmsg.ErrPageNotFound, nil)
	})
	pprof.Register(g)
	//webHandler.GET("/metrics", gin.WrapH(promhttp.Handler()))
	g.GET("/health", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})
	return nil
}

// installRouters 安装 migaudu-bridge 接口路由.
func installRouters(g *gin.Engine) error {

	addyAliases := g.Group("/api/v1/aliases").POST("")

	simplelogin := g.Group("/api/alias/random/new").POST("")

	return nil
}

type AddyAliases struct {
	Domain      string `json:"domain"`
	Description string `json:"description"`
}

type Simplelogin struct {
	Note string `json:"note"`
}
