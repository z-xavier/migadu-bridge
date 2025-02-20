package log

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
	"sync/atomic"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	slogzerolog "github.com/samber/slog-zerolog/v2"

	"migadu-bridge/internal/pkg/common"
)

// Logger 定义了项目的日志接口. 该接口只包含了支持的日志记录方法.
type Logger interface {
	C(ctx context.Context) Logger
	WithError(err error) Logger
	WithField(key string, value any) Logger
	Debug(a ...any)
	Debugw(msg string, keysAndValues ...any)
	Debugf(format string, args ...any)
	Info(a ...any)
	Infow(msg string, keysAndValues ...any)
	Infof(format string, args ...any)
	Warn(a ...any)
	Warnw(msg string, keysAndValues ...any)
	Warnf(format string, args ...any)
	Error(a ...any)
	Errorw(msg string, keysAndValues ...any)
	Errorf(format string, args ...any)
	Panic(a ...any)
	Panicw(msg string, keysAndValues ...any)
	Panicf(format string, args ...any)
	Fatal(a ...any)
	Fatalw(msg string, keysAndValues ...any)
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
	std atomic.Pointer[sLogger]
)

var (
	TraceIdKey     = common.XRequestIDKey
	ErrorFieldName = zerolog.ErrorFieldName
)

// 初始化时设置默认 Logger
func init() {
	std.Store(NewLogger(NewOptions()))
}

// Init 使用指定的选项初始化 Logger.
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std.Store(NewLogger(opts)) // 原子存储
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
	if !opts.DisableStacktrace {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}

	zerolog.InterfaceMarshalFunc = func(v interface{}) ([]byte, error) {
		var buf bytes.Buffer
		encoder := sonic.ConfigDefault.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(v)
		if err != nil {
			return nil, err
		}
		b := buf.Bytes()
		if len(b) > 0 {
			// Remove trailing \n which is added by Encode.
			return b[:len(b)-1], nil
		}
		return b, nil
	}

	// 配置输出 Writer
	var writers []io.Writer

	// 设置日志显示格式
	if opts.Format == "console" {
		writers = append(writers, zerolog.ConsoleWriter{
			Out: os.Stderr,
		})
	} else {
		writers = append(writers, os.Stderr)
	}

	// 设置日志输出目录
	// 文件输出
	if opts.FilePath != "" {
		file, _ := os.OpenFile(opts.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		writers = append(writers, file)
	}

	zerologLogger := zerolog.New(zerolog.MultiLevelWriter(writers...))

	requestIdFormCtx := func(ctx context.Context) []slog.Attr {
		if requestID := ctx.Value(TraceIdKey); requestID != nil {
			if reqID, ok := requestID.(string); ok {
				return []slog.Attr{slog.String(TraceIdKey, reqID)}
			}
		}
		return []slog.Attr{}
	}

	// 使用 cfg 创建 *slog.Logger 对象
	s := slog.New(slogzerolog.Option{
		Level:           sLevel,
		Logger:          &zerologLogger,
		AttrFromContext: []func(ctx context.Context) []slog.Attr{requestIdFormCtx},
		AddSource:       !opts.DisableSource,
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
		s: l.s.With(ErrorFieldName, err),
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

func (l *sLogger) Debugw(msg string, keysAndValues ...any) {
	l.s.DebugContext(l.context, msg, keysAndValues...)
}

func (l *sLogger) Debugf(format string, args ...any) {
	l.s.DebugContext(l.context, fmt.Sprintf(format, args...))
}

func (l *sLogger) Info(a ...any) {
	l.s.InfoContext(l.context, fmt.Sprint(a...))
}

func (l *sLogger) Infow(msg string, keysAndValues ...any) {
	l.s.InfoContext(l.context, msg, keysAndValues...)
}

func (l *sLogger) Infof(format string, args ...any) {
	l.s.InfoContext(l.context, fmt.Sprintf(format, args...))
}

func (l *sLogger) Warn(a ...any) {
	l.s.WarnContext(l.context, fmt.Sprint(a...))
}

func (l *sLogger) Warnw(msg string, keysAndValues ...any) {
	l.s.WarnContext(l.context, msg, keysAndValues...)
}

func (l *sLogger) Warnf(format string, args ...any) {
	l.s.WarnContext(l.context, fmt.Sprintf(format, args...))
}

func (l *sLogger) Error(a ...any) {
	l.s.ErrorContext(l.context, fmt.Sprint(a...))
}

func (l *sLogger) Errorw(msg string, keysAndValues ...any) {
	l.s.ErrorContext(l.context, msg, keysAndValues...)
}

func (l *sLogger) Errorf(format string, args ...any) {
	l.s.ErrorContext(l.context, fmt.Sprintf(format, args...))
}

func (l *sLogger) Panic(a ...any) {
	msg := fmt.Sprint(a...)
	l.s.ErrorContext(l.context, msg)
	panic(msg)
}

func (l *sLogger) Panicw(msg string, keysAndValues ...any) {
	l.s.ErrorContext(l.context, msg, keysAndValues...)
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

func (l *sLogger) Fatalw(msg string, keysAndValues ...any) {
	l.s.ErrorContext(l.context, msg, keysAndValues...)
	os.Exit(1)
}

func (l *sLogger) Fatalf(format string, a ...any) {
	l.s.ErrorContext(l.context, fmt.Sprintf(format, a...))
	os.Exit(1)
}
