package log

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"migadu-bridge/internal/pkg/common"
)

// Logger 定义了项目的日志接口. 该接口只包含了支持的日志记录方法.
type Logger interface {
	Debugf(format string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infof(format string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnf(format string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorf(format string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicf(format string, args ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync()
}

type sLogger struct {
	s *slog.Logger
}

var _ Logger = &sLogger{}

var (
	mu sync.Mutex

	// std 定义了默认的全局 Logger.
	std = NewLogger(NewOptions())
)

// Init 使用指定的选项初始化 Logger.
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

// NewLogger 根据传入的 opts 创建 Logger.
func NewLogger(opts *Options) *sLogger {
	if opts == nil {
		opts = NewOptions()
	}

	// 将文本格式的日志级别，例如 info 转换为 slog.Level 类型以供后面使用
	var sLevel slog.Level
	if err := sLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		// 如果指定了非法的日志级别，则默认使用 info 级别
		sLevel = slog.LevelInfo
	}

	// 使用 cfg 创建 *slog.Logger 对象
	s := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     sLevel,
	}))
	slog.SetDefault(s)

	logger := &sLogger{s}

	return logger
}

// Sync 调用底层 zap.Logger 的 Sync 方法，将缓存中的日志刷新到磁盘文件中. 主程序需要在退出前调用 Sync.
func Sync() { std.Sync() }

func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}

func Debugf(template string, args ...interface{}) {
	std.z.Sugar().Debugf(template, args...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.z.Sugar().Debugf(format, args...)
}

// Debugw 输出 debug 级别的日志.
func Debugw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Debugw(msg, keysAndValues...)
}

func Infof(template string, args ...interface{}) {
	std.z.Sugar().Infof(template, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.z.Sugar().Infof(format, args...)
}

// Infow 输出 info 级别的日志.
func Infow(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Infow(msg, keysAndValues...)
}

func Warnf(template string, args ...interface{}) {
	std.z.Sugar().Warnf(template, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.z.Sugar().Warnf(format, args...)
}

// Warnw 输出 warning 级别的日志.
func Warnw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Warnw(msg, keysAndValues...)
}

func Errorf(template string, args ...interface{}) {
	std.z.Sugar().Errorf(template, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.z.Sugar().Errorf(format, args...)
}

// Errorw 输出 error 级别的日志.
func Errorw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Errorw(msg, keysAndValues...)
}

func Panicf(template string, args ...interface{}) {
	std.z.Sugar().Panicf(template, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.z.Sugar().Panicf(format, args...)
}

// Panicw 输出 panic 级别的日志.
func Panicw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Panicw(msg, keysAndValues...)
}

func Fatalf(template string, args ...interface{}) {
	std.z.Sugar().Fatalf(template, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.z.Sugar().Fatalf(format, args...)
}

// Fatalw 输出 fatal 级别的日志.
func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Fatalw(msg, keysAndValues...)
}

// C 解析传入的 context，尝试提取关注的键值，并添加到 zap.Logger 结构化日志中.
func C(ctx context.Context) *sLogger {
	return std.C(ctx)
}

func (l *sLogger) C(ctx context.Context) *sLogger {
	if requestID := ctx.Value(common.XRequestIDKey); requestID != nil {
		lc.z = lc.z.With(zap.Any(common.XRequestIDKey, requestID))
	}
	return l.s.WarnContext()
}

// clone 深度拷贝 zapLogger.
func (l *zapLogger) clone() *zapLogger {
	lc := *l
	return &lc
}
