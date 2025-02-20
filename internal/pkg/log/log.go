package log

import (
	"context"
)

func C(ctx context.Context) Logger {
	return std.C(ctx)
}

func WithError(err error) Logger {
	return std.WithError(err)
}

func WithField(key string, value any) Logger {
	return std.WithField(key, value)
}

func Debug(a ...any) {
	std.Debug(a...)
}

func Debugf(format string, args ...any) {
	std.Debugf(format, args...)
}

func Info(a ...any) {
	std.Info(a...)
}

func Infof(format string, args ...any) {
	std.Infof(format, args...)
}

func Warn(a ...any) {
	std.Warn(a...)
}

func Warnf(format string, args ...any) {
	std.Warnf(format, args...)
}

func Error(a ...any) {
	std.Error(a...)
}

func Errorf(format string, args ...any) {
	std.Errorf(format, args...)
}

func Panic(a ...any) {
	std.Panic(a...)
}

func Panicf(format string, a ...any) {
	std.Panicf(format, a...)
}

func Fatal(a ...any) {
	std.Fatal(a...)
}

func Fatalf(format string, a ...any) {
	std.Fatalf(format, a...)
}
