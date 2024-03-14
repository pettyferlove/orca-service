package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync/atomic"
)

type Logger struct {
	*zap.Logger
}

var defaultLogger atomic.Value

func init() {
	defaultLogger.Store(NewLogger(InfoLevel))
}

// NewLogger 创建一个新的日志对象
func NewLogger(level string) *Logger {
	var zapLevel zapcore.Level
	switch level {
	case DebugLevel:
		zapLevel = zapcore.DebugLevel
	case InfoLevel:
		zapLevel = zapcore.InfoLevel
	case WarnLevel:
		zapLevel = zapcore.WarnLevel
	case ErrorLevel:
		zapLevel = zapcore.ErrorLevel
	case FatalLevel:
		zapLevel = zapcore.FatalLevel
	case PanicLevel:
		zapLevel = zapcore.PanicLevel
	case TraceLevel:
		zapLevel = zapcore.DebugLevel
	default:
		zapLevel = zapcore.InfoLevel
	}
	coreList := make([]zapcore.Core, 0)
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 修改这行来改变时间格式
	encoder := zapcore.NewConsoleEncoder(config)
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapLevel))
	core := zapcore.NewTee(coreList...)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	return &Logger{
		Logger: zapLogger,
	}
}

// SetDefaultLoggerLevel 设置默认日志级别
func SetDefaultLoggerLevel(level string) {
	defaultLogger.Store(NewLogger(level))
}

// DefaultLogger 获取默认日志对象
func DefaultLogger() *Logger {
	return defaultLogger.Load().(*Logger)
}

func (l *Logger) Printf(s string, i ...interface{}) {
	defaultLogger.Load().(*Logger).Info(fmt.Sprintf(s, i...))
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
