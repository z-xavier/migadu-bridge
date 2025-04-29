package log

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

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

const defaultDepth int = 4

type sLogger struct {
	context context.Context
	// depth slog 官方不支持调整 caller 的 skip，给出了下列的解决方案。
	// https://github.com/golang/go/issues/59145#issuecomment-1481920720
	depth int
	s     *slog.Logger
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

// writerFactory 用于创建不同类型的日志输出writer
type writerFactory struct {
	format string
}

// createWriter 根据输出类型创建对应的writer
func (f *writerFactory) createWriter(output string) (io.Writer, error) {
	switch output {
	case "stdout", "stderr":
		return f.createStdWriter(output)
	default:
		return f.createFileWriter(output)
	}
}

// createStdWriter 创建标准输出的writer
func (f *writerFactory) createStdWriter(output string) (io.Writer, error) {
	stdOut := os.Stdout
	if output == "stderr" {
		stdOut = os.Stderr
	}

	if f.format == "console" {
		return zerolog.ConsoleWriter{Out: stdOut}, nil
	}
	return stdOut, nil
}

// createFileWriter 创建文件输出的writer
func (f *writerFactory) createFileWriter(path string) (io.Writer, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file %q: %v", path, err)
	}

	if f.format == "console" {
		return zerolog.ConsoleWriter{Out: file}, nil
	}
	return file, nil
}

// NewLogger 根据传入的 opts 创建 Logger.
func NewLogger(opts *Options) *sLogger {
	if opts == nil {
		opts = NewOptions()
	}

	// 设置日志级别
	var sLevel slog.Level
	if err := sLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		sLevel = slog.LevelInfo
	}

	// 设置错误跟踪
	if !opts.DisableStacktrace {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}

	// 配置 JSON 序列化
	zerolog.InterfaceMarshalFunc = func(v any) ([]byte, error) {
		var buf bytes.Buffer
		encoder := sonic.ConfigDefault.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(v)
		if err != nil {
			return nil, err
		}
		b := buf.Bytes()
		if len(b) > 0 {
			return b[:len(b)-1], nil // Remove trailing \n
		}
		return b, nil
	}

	// 配置输出
	if len(opts.OutputPaths) == 0 {
		opts.OutputPaths = []string{"stdout"}
	}

	factory := &writerFactory{format: opts.Format}
	var writers []io.Writer

	for _, output := range opts.OutputPaths {
		writer, err := factory.createWriter(output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		writers = append(writers, writer)
	}

	// 创建 zerolog logger
	zerologLogger := zerolog.New(zerolog.MultiLevelWriter(writers...))

	requestIdFormCtx := func(ctx context.Context) []slog.Attr {
		if requestID := ctx.Value(TraceIdKey); requestID != nil {
			if reqID, ok := requestID.(string); ok {
				return []slog.Attr{slog.String(TraceIdKey, reqID)}
			}
		}
		return []slog.Attr{}
	}

	replaceSource := func(groups []string, a slog.Attr) slog.Attr {
		// TODO 等 zerolog 官方适配 slog，后续观察
		// 当前 zerlog 不支持 group，故第三方 zerolog - slog 适配器，source 是分三个字段处理的。
		// 无法优雅适配。如果选用 slog 后端，可以这样处理。
		// https://github.com/golang/go/issues/59145#issuecomment-1481920720
		//if a.Key == slog.SourceKey {
		//	// 调整skip值以匹配你的调用层级
		//	const skipOffset = 8 // 可能需要根据实际情况调整此值
		//	pc := make([]uintptr, 1)
		//	if n := runtime.Callers(skipOffset, pc); n > 0 {
		//		frame, _ := runtime.CallersFrames(pc).Next()
		//		return slog.Any(slog.SourceKey, &slog.Source{
		//			Function: frame.Function,
		//			File:     frame.File,
		//			Line:     frame.Line,
		//		})
		//	}
		//}
		return a
	}

	// 使用 cfg 创建 *slog.Logger 对象
	s := slog.New(slogzerolog.Option{
		Level:           sLevel,
		Logger:          &zerologLogger,
		AttrFromContext: []func(ctx context.Context) []slog.Attr{requestIdFormCtx},
		AddSource:       !opts.DisableSource,
		ReplaceAttr:     replaceSource,
	}.NewZerologHandler())

	slog.SetDefault(s)

	logger := &sLogger{
		depth: defaultDepth + 1,
		s:     s,
	}

	return logger
}

func (l *sLogger) log(level slog.Level, msg string, keysAndValues ...any) {
	var pcs [1]uintptr
	runtime.Callers(l.depth, pcs[:])
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(keysAndValues...)

	if l.context == nil {
		l.context = context.Background()
	}

	_ = l.s.Handler().Handle(l.context, r)
}

func (l *sLogger) C(ctx context.Context) Logger {
	return &sLogger{
		context: ctx,
		depth:   defaultDepth,
		s:       l.s,
	}
}

func (l *sLogger) WithError(err error) Logger {
	return &sLogger{
		context: l.context,
		depth:   defaultDepth,
		s:       l.s.With(ErrorFieldName, err),
	}
}

func (l *sLogger) WithField(key string, value any) Logger {
	return &sLogger{
		context: l.context,
		depth:   defaultDepth,
		s:       l.s.With(key, value),
	}
}

func (l *sLogger) Log(level slog.Level, a ...any) {
	if !l.s.Enabled(l.context, level) {
		return
	}
	l.log(level, fmt.Sprint(a...))
}

func (l *sLogger) Logw(level slog.Level, msg string, keysAndValues ...any) {
	if !l.s.Enabled(l.context, level) {
		return
	}
	l.log(level, msg, keysAndValues...)
}

func (l *sLogger) Logf(level slog.Level, format string, a ...any) {
	if !l.s.Enabled(l.context, level) {
		return
	}
	l.log(level, fmt.Sprintf(format, a...))
}

func (l *sLogger) Debug(a ...any) {
	l.Log(slog.LevelDebug, a...)
}

func (l *sLogger) Debugw(msg string, keysAndValues ...any) {
	l.Logw(slog.LevelDebug, msg, keysAndValues...)
}

func (l *sLogger) Debugf(format string, a ...any) {
	l.Logf(slog.LevelDebug, format, a...)
}

func (l *sLogger) Info(a ...any) {
	l.Log(slog.LevelInfo, a...)
}

func (l *sLogger) Infow(msg string, keysAndValues ...any) {
	l.Logw(slog.LevelInfo, msg, keysAndValues...)
}

func (l *sLogger) Infof(format string, a ...any) {
	l.Logf(slog.LevelInfo, format, a...)
}

func (l *sLogger) Warn(a ...any) {
	l.Log(slog.LevelWarn, a...)
}

func (l *sLogger) Warnw(msg string, keysAndValues ...any) {
	l.Logw(slog.LevelWarn, msg, keysAndValues...)
}

func (l *sLogger) Warnf(format string, a ...any) {
	l.Logf(slog.LevelWarn, format, a...)
}

func (l *sLogger) Error(a ...any) {
	l.Log(slog.LevelError, a...)
}

func (l *sLogger) Errorw(msg string, keysAndValues ...any) {
	l.Logw(slog.LevelError, msg, keysAndValues...)
}

func (l *sLogger) Errorf(format string, a ...any) {
	l.Logf(slog.LevelError, format, a...)
}

func (l *sLogger) Panic(a ...any) {
	l.Log(slog.LevelError, a...)
	panic(fmt.Sprint(a...))
}

func (l *sLogger) Panicw(msg string, keysAndValues ...any) {
	l.Logw(slog.LevelError, msg, keysAndValues...)
	panic(msg)
}

func (l *sLogger) Panicf(format string, a ...any) {
	l.Logf(slog.LevelError, format, a...)
	panic(fmt.Sprintf(format, a...))
}

func (l *sLogger) Fatal(a ...any) {
	l.Log(slog.LevelError, a...)
	os.Exit(1)
}

func (l *sLogger) Fatalw(msg string, keysAndValues ...any) {
	l.Logw(slog.LevelError, msg, keysAndValues...)
	os.Exit(1)
}

func (l *sLogger) Fatalf(format string, a ...any) {
	l.Logf(slog.LevelError, format, a...)
	os.Exit(1)
}
