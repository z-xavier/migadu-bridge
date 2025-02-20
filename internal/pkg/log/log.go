package log

import (
	"context"
)

func C(ctx context.Context) Logger {
	return std.Load().C(ctx)
}

func WithError(err error) Logger {
	return std.Load().WithError(err)
}

func WithField(key string, value any) Logger {
	return std.Load().WithField(key, value)
}

func Debug(a ...any) {
	std.Load().Debug(a...)
}

func Debugw(msg string, keysAndValues ...any) {
	std.Load().Debugw(msg, keysAndValues...)
}

func Debugf(format string, args ...any) {
	std.Load().Debugf(format, args...)
}

func Info(a ...any) {
	std.Load().Info(a...)
}

func Infow(msg string, keysAndValues ...any) {
	std.Load().Infow(msg, keysAndValues...)
}

func Infof(format string, args ...any) {
	std.Load().Infof(format, args...)
}

func Warn(a ...any) {
	std.Load().Warn(a...)
}

func Warnw(msg string, keysAndValues ...any) {
	std.Load().Warnw(msg, keysAndValues...)
}

func Warnf(format string, args ...any) {
	std.Load().Warnf(format, args...)
}

func Error(a ...any) {
	std.Load().Error(a...)
}

func Errorw(msg string, keysAndValues ...any) {
	std.Load().Errorw(msg, keysAndValues...)
}

func Errorf(format string, args ...any) {
	std.Load().Errorf(format, args...)
}

func Panic(a ...any) {
	std.Load().Panic(a...)
}

func Panicw(msg string, keysAndValues ...any) {
	std.Load().Panicw(msg, keysAndValues...)
}

func Panicf(format string, a ...any) {
	std.Load().Panicf(format, a...)
}

func Fatal(a ...any) {
	std.Load().Fatal(a...)
}

func Fatalw(msg string, keysAndValues ...any) {
	std.Load().Fatalw(msg, keysAndValues...)
}

func Fatalf(format string, a ...any) {
	std.Load().Fatalf(format, a...)
}
