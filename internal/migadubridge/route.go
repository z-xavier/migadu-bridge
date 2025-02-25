package migadubridge

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/controller/aliases"
	"migadu-bridge/internal/migadubridge/controller/call_logs"
	"migadu-bridge/internal/migadubridge/controller/tokens"
	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/core"
	"migadu-bridge/internal/pkg/errmsg"
	"migadu-bridge/internal/pkg/log"
)

// installInteriorWebRouters Gin 内部服务器
func installInteriorWebRouters(g *gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errmsg.ErrPageNotFound, nil)
	})

	pprof.Register(g)
	g.GET("/health", func(c *gin.Context) {
		log.C(c).Info("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	tc := tokens.New(store.S)
	cc := call_logs.New(store.S)
	ac := aliases.New(store.S)

	v1 := g.Group("/api/v1")
	{
		tokenV1 := v1.Group("/tokens")
		{
			tokenV1.POST("", tc.Create)
			tokenV1.DELETE(":tokenId", tc.Delete)
			tokenV1.PUT(":tokenId", tc.Put)
			tokenV1.PATCH(":tokenId", tc.Patch)
			tokenV1.GET("", tc.List)
			tokenV1.GET(":tokenId", tc.Get)
		}

		callLogV1 := v1.Group("/calllogs")
		{
			callLogV1.GET("", cc.List)
		}

		aliasV1 := v1.Group("/aliases")
		{
			aliasV1.GET("", ac.List)
		}
	}

	return nil
}

// installRouters 安装 migaudu-provider 接口路由.
func installRouters(g *gin.Engine) error {
	// TODO
	// V1 版本先直接根据 path 区分路由
	// V2 版本中 如果相同 path 根据 Token 获取对应的控制器

	//addyAliases := g.Group("/api/v1/aliases").POST("")
	//
	//simplelogin := g.Group("/api/alias/random/new").POST("")

	return nil
}
