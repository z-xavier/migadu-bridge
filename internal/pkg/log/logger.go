package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	slogzerolog "github.com/samber/slog-zerolog/v2"

	"migadu-bridge/internal/pkg/common"
)

type Fields map[string]any

// Logger 定义了项目的日志接口. 该接口只包含了支持的日志记录方法.
type Logger interface {
	C(ctx context.Context) Logger
	WithError(err error) Logger
	WithField(key string, value any) Logger
	Debug(a ...any)
	Debugf(format string, args ...any)
	Info(a ...any)
	Infof(format string, args ...any)
	Warn(a ...any)
	Warnf(format string, args ...any)
	Error(a ...any)
	Errorf(format string, args ...any)
	Panic(a ...any)
	Panicf(format string, args ...any)
	Fatal(a ...any)
	Fatalf(format string, args ...any)
}

type sLogger struct {
	context context.Context
	s       *slog.Logger
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

	// 设置错误跟踪
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerologLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

	requestIdFormCtx := func(ctx context.Context) []slog.Attr {
		if requestID := ctx.Value(common.XRequestIDKey); requestID != nil {
			if reqID, ok := requestID.(string); ok {
				return []slog.Attr{slog.String(common.XRequestIDKey, reqID)}
			}
		}
		return []slog.Attr{}
	}

	// 使用 cfg 创建 *slog.Logger 对象
	s := slog.New(slogzerolog.Option{
		Level:           sLevel,
		Logger:          &zerologLogger,
		AttrFromContext: []func(ctx context.Context) []slog.Attr{requestIdFormCtx},
		AddSource:       !opts.DisableCaller,
	}.NewZerologHandler())

	slog.SetDefault(s)

	logger := &sLogger{
		s: s,
	}

	return logger
}

func (l *sLogger) C(ctx context.Context) Logger {
	return &sLogger{
		context: ctx,
		s:       l.s,
	}
}

func (l *sLogger) WithError(err error) Logger {
	return &sLogger{
		s: l.s.With(zerolog.ErrorFieldName, err),
	}
}

func (l *sLogger) WithField(key string, value any) Logger {
	return &sLogger{
		context: l.context,
		s:       l.s.With(key, value),
	}
}

func (l *sLogger) Debug(a ...any) {
	l.s.DebugContext(l.context, fmt.Sprint(a...))
}

func (l *sLogger) Debugf(format string, args ...any) {
	l.s.DebugContext(l.context, fmt.Sprintf(format, args...))
}

func (l *sLogger) Info(a ...any) {
	l.s.InfoContext(l.context, fmt.Sprint(a...))
}

func (l *sLogger) Infof(format string, args ...any) {
	l.s.InfoContext(l.context, fmt.Sprintf(format, args...))
}

func (l *sLogger) Warn(a ...any) {
	l.s.WarnContext(l.context, fmt.Sprint(a...))
}

func (l *sLogger) Warnf(format string, args ...any) {
	l.s.WarnContext(l.context, fmt.Sprintf(format, args...))
}

func (l *sLogger) Error(a ...any) {
	l.s.ErrorContext(l.context, fmt.Sprint(a...))
}

func (l *sLogger) Errorf(format string, args ...any) {
	l.s.ErrorContext(l.context, fmt.Sprintf(format, args...))
}

func (l *sLogger) Panic(a ...any) {
	msg := fmt.Sprint(a...)
	l.s.ErrorContext(l.context, msg)
	panic(msg)
}

func (l *sLogger) Panicf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	l.s.ErrorContext(l.context, msg)
	panic(msg)
}

func (l *sLogger) Fatal(a ...any) {
	l.s.ErrorContext(l.context, fmt.Sprint(a...))
	os.Exit(1)
}

func (l *sLogger) Fatalf(format string, a ...any) {
	l.s.ErrorContext(l.context, fmt.Sprintf(format, a...))
	os.Exit(1)
}
