package migadubridge

import (
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "migadu-bridge/docs"
	"migadu-bridge/internal/migadubridge/controller/aliases"
	"migadu-bridge/internal/migadubridge/controller/bridges"
	"migadu-bridge/internal/migadubridge/controller/call_logs"
	"migadu-bridge/internal/migadubridge/controller/tokens"
	mstatic "migadu-bridge/internal/migadubridge/static"
	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/core"
	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/pkg/pprof"
)

var indexContent = sync.OnceValue(func() []byte {
	// 在程序启动时读取并缓存 index.html
	fs, err := mstatic.GetFS()
	if err != nil {
		log.WithError(err).Error("failed to get static file system")
		return nil
	}

	indexFile, err := fs.Open("index.html")
	if err != nil {
		log.WithError(err).Error("failed to open index.html")
		return nil
	}

	defer func(indexFile http.File) {
		err := indexFile.Close()
		if err != nil {
			log.WithError(err).Warn("failed to close index.html")
		}
	}(indexFile)

	content, err := io.ReadAll(indexFile)
	if err != nil {
		log.WithError(err).Error("failed to read index.html")
		return nil
	}
	return content
})

const (
	HttpAcceptHeader    = "Accept"
	ContentTypeText     = "text/html"
	ContentTypeTextFull = "text/html; charset=utf-8"
)

// installInteriorWebRouters Gin 内部服务器
func installInteriorWebRouters(g *gin.Engine) error {
	fs, err := mstatic.GetFS()
	if err != nil {
		log.WithError(err).Error("failed to get static file system")
		return nil
	}
	g.Use(static.Serve("/", fs))

	// 处理所有未匹配的路由，返回 index.html
	g.NoRoute(func(c *gin.Context) {
		// 只对 HTML 请求返回 index.html
		if strings.Contains(c.Request.Header.Get(HttpAcceptHeader), ContentTypeText) {
			c.Data(http.StatusOK, ContentTypeTextFull, indexContent())
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	g.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	pprof.Register(g)
	g.GET("/health", func(c *gin.Context) {
		log.C(c).Info("Healthz function called")
		c.String(200, "ok")
	})

	// Swagger documentation
	//docs.SwaggerInfo.BasePath = "/api/v1"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tc := tokens.New(store.S)
	cc := call_logs.New(store.S)
	ac := aliases.New(store.S)

	v1 := g.Group("/api/v1")
	{
		tokenV1 := v1.Group("/tokens")
		{
			tokenV1.POST("", core.HandleResult(tc.Create))
			tokenV1.DELETE(":tokenId", core.HandleResult(tc.Delete))
			tokenV1.PUT(":tokenId", core.HandleResult(tc.Put))
			tokenV1.PATCH(":tokenId", core.HandleResult(tc.Patch))
			tokenV1.GET("", core.HandleResult(tc.List))
			tokenV1.GET(":tokenId", core.HandleResult(tc.Get))
		}

		callLogV1 := v1.Group("/calllogs")
		{
			callLogV1.GET("", core.HandleResult(cc.List))
		}

		aliasV1 := v1.Group("/aliases")
		{
			aliasV1.GET("", core.HandleResult(ac.List))
		}
	}

	return nil
}

// installRouters 安装 migaudu-provider 接口路由.
func installRouters(g *gin.Engine) error {
	// V1 版本先直接根据 path 区分路由
	// V2 版本中 如果相同 path 根据 Token 获取对应的控制器

	b := bridges.New(store.S)

	// addyAliases
	g.Group("/api/v1/aliases").POST("", b.AddyAliases)

	// simplelogin
	g.Group("/api/alias/random/new").POST("", b.SLAliasRandomNew)

	return nil
}
