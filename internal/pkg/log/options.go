package log

import (
	"log/slog"
)

// Options 包含与日志相关的配置项.
type Options struct {
	// 是否开启 caller，如果开启会在日志中显示调用日志所在的文件和行号
	DisableCaller bool
	// 指定日志级别，可选值：debug, info, warn, error, panic, fatal
	Level string
}

// NewOptions 创建一个带有默认参数的 Options 对象.
func NewOptions() *Options {
	return &Options{
		DisableCaller: false,
		Level:         slog.LevelInfo.String(),
	}
}
