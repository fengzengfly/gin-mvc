package log

import (
	"context"
	"fmt"
	"gin-mvc/pkg/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var Logger *zap.Logger

// InitLogger 初始化日志
func InitLogger(logConfig *config.LogConfig) error {

	// 读取日志配置
	if err := viper.UnmarshalKey("log", &logConfig); err != nil {
		return fmt.Errorf("解析日志配置失败: %v", err)
	}

	// 创建日志目录
	logDir := filepath.Dir(logConfig.OutputPaths[0])
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return fmt.Errorf("创建日志目录失败: %v", err)
	}

	// 日志轮转
	writeSyncer := getLogWriter(
		logConfig.OutputPaths[0],
		logConfig.MaxSize,
		logConfig.MaxBackups,
		logConfig.MaxAge,
	)

	// 日志编码配置
	encoder := getEncoder(logConfig.Format)

	// 日志级别
	level := getLogLevel(logConfig.Level)

	// 创建核心日志记录器
	core := zapcore.NewCore(
		encoder,
		writeSyncer,
		level,
	)

	// 创建日志记录器
	Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	return nil
}

// 获取日志写入器
func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 获取日志编码器
func getEncoder(format string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 自定义时间格式
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	if format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 获取日志级别
func getLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

// Debug 日志扩展方法
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

// 获取调用者信息
func getCallerInfo() (string, int, string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "", 0, ""
	}

	funcName := runtime.FuncForPC(pc).Name()
	return file, line, funcName
}

// WithContext 上下文日志记录
func WithContext(ctx context.Context) *zap.Logger {
	// 可以从 context 中提取 traceID 等信息
	traceID := ctx.Value("trace_id")
	if traceID != nil {
		return Logger.With(zap.Any("trace_id", traceID))
	}
	return Logger
}
