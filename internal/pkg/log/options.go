package log

import (
	"log/slog"
)

// Options 包含与日志相关的配置项.
type Options struct {
	// 是否开启 source，如果开启会在日志中显示调用日志所在的文件和行号
	DisableSource bool
	// 是否禁止在 panic 及以上级别打印堆栈信息
	DisableStacktrace bool
	// 指定日志级别，可选值：debug, info, warn, error, panic, fatal
	Level string
	// 指定日志显示格式，可选值：console, json
	Format string
	// 指定日志输出位置
	OutputPaths []string
}

// NewOptions 创建一个带有默认参数的 Options 对象.
func NewOptions() *Options {
	return &Options{
		DisableSource:     false,
		DisableStacktrace: false,
		Level:             slog.LevelInfo.String(),
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}
