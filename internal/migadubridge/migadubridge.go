package migadubridge

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"migadu-bridge/internal/pkg/config"
	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/internal/pkg/middleware"
)

var cfgFile string

// NewMigaduBridgeCommand 创建一个 *cobra.Command 对象. 之后，可以使用 Command 对象的 Execute 方法来启动应用程序.
func NewMigaduBridgeCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "migadu-provider",
		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := mbInit(); err != nil {
				return err
			}
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().StringVarP(&cfgFile, "config", "c", "", "The path to the migadu-provider configuration file. Empty string for no configuration file.")
	return cmd
}

func mbInit() error {
	if err := config.InitConfig(cfgFile); err != nil {
		return err
	}
	log.Init(config.LogOptions())
	return nil
}

// run 函数是实际的业务代码入口函数.
func run() error {

	// 设置 Gin 模式
	gin.SetMode(config.GetConfig().ServerConf.RunMode)

	// 初始化 store 层
	if err := initStore(); err != nil {
		return err
	}

	// 创建 Gin 内部服务器
	interiorWebHandler := gin.New()
	interiorWebHandler.Use(gin.Recovery(),
		middleware.Cors(),
		middleware.RequestId(),
		middleware.GinLog(),
		middleware.ResponseTime(),
	)
	if err := installInteriorWebRouters(interiorWebHandler); err != nil {
		return err
	}

	// 创建 Gin 引擎
	webHandler := gin.New()
	webHandler.Use(gin.Recovery(),
		middleware.Cors(),
		middleware.RequestId(),
		middleware.GinLog(),
		middleware.ResponseTime(),
	)
	if err := installRouters(webHandler); err != nil {
		return err
	}

	// 创建并运行内部 HTTP 服务器
	interiorWebSrv := startInsecureServer(config.GetConfig().ServerConf.InteriorWebAddr, interiorWebHandler)

	// 创建并运行 HTTP 服务器
	webSrv := startInsecureServer(config.GetConfig().ServerConf.WebAddr, webHandler)

	// 等待中断信号优雅地关闭服务器（10 秒超时)。
	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Info("Shutting down server ...")

	// 创建 ctx 用于通知服务器 goroutine, 它有 10 秒时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 10 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 10 秒就超时退出
	if err := webSrv.Shutdown(ctx); err != nil {
		log.WithError(err).Error("Web Server forced to shutdown")
		return err
	}
	if err := interiorWebSrv.Shutdown(ctx); err != nil {
		log.WithError(err).Error("InteriorWeb Server forced to shutdown")
		return err
	}
	log.Info("Server exiting")
	return nil
}

// startInsecureServer 创建并运行 HTTP 服务器.
func startInsecureServer(addr string, g *gin.Engine) *http.Server {
	// 创建 HTTP Server 实例
	httpsrv := &http.Server{Addr: addr, Handler: g}
	// 运行 HTTP 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	// 打印一条日志，用来提示 HTTP 服务已经起来，方便排障
	log.WithField("addr", addr).Infow("Start to listening the incoming requests on http address")
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Fatal()
		}
	}()
	return httpsrv
}
