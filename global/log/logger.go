package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync/atomic"
)

type Logger struct {
	*zap.Logger
}

var defaultLogger atomic.Value

func init() {
	defaultLogger.Store(NewLogger())
}

func NewLogger() *Logger {
	coreList := make([]zapcore.Core, 0)
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))

	debugJackLogger := &lumberjack.Logger{
		Filename:   "./logs/debug.log",
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 30,
		LocalTime:  true,
		Compress:   false,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(debugJackLogger), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})))

	infoJackLogger := &lumberjack.Logger{
		Filename:   "./logs/info.log",
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 30,
		LocalTime:  true,
		Compress:   false,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(infoJackLogger), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})))

	warnJackLogger := &lumberjack.Logger{
		Filename:   "./logs/warn.log",
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 30,
		LocalTime:  true,
		Compress:   false,
	}

	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(warnJackLogger), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})))

	errorJackLogger := &lumberjack.Logger{
		Filename:   "./logs/error.log",
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 30,
		LocalTime:  true,
		Compress:   false,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(errorJackLogger), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})))

	core := zapcore.NewTee(coreList...)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{
		Logger: zapLogger,
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.Logger.Debug(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.Logger.Warn(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.Logger.Fatal(fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(format string, v ...interface{}) {
	l.Logger.Panic(fmt.Sprintf(format, v...))
}

func Debug(format string, v ...interface{}) {
	defaultLogger.Load().(*Logger).Debug(fmt.Sprintf(format, v...))
}

func Info(format string, v ...interface{}) {
	defaultLogger.Load().(*Logger).Info(fmt.Sprintf(format, v...))
}

func Warn(format string, v ...interface{}) {
	defaultLogger.Load().(*Logger).Warn(fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	defaultLogger.Load().(*Logger).Error(fmt.Sprintf(format, v...))
}

func Fatal(format string, v ...interface{}) {
	defaultLogger.Load().(*Logger).Error(fmt.Sprintf(format, v...))
}

func Panic(format string, v ...interface{}) {
	defaultLogger.Load().(*Logger).Error(fmt.Sprintf(format, v...))
}
