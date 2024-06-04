package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func InitLogger(appName, logPath string, debugMode bool) error {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	logWriter, err := rotatelogs.New(
		logPath+string(filepath.Separator)+appName+".log.%Y_%m_%d_%H",
		rotatelogs.WithLinkName(logPath+string(filepath.Separator)+appName+".log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(time.Minute),
	)
	if err != nil {
		return fmt.Errorf("failed to create rotatelogs: %w", err)
	}

	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(logWriter), zapcore.InfoLevel)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)

	if debugMode {
		core := zapcore.NewTee(fileCore, consoleCore)
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
		return nil
	}

	logger = zap.New(fileCore, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	return nil
}

func Debugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

func Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

func Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

func DPanicf(msg string, args ...interface{}) {
	logger.DPanicf(msg, args...)
}

func Panicf(msg string, args ...interface{}) {
	logger.Panicf(msg, args...)
}

func Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args...)
}
