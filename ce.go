package ce

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// code commit hash
var Version string

var DefaultLogger *zap.Logger
var DefaultAtomicLevel zap.AtomicLevel

func init() {
	DefaultAtomicLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	DefaultLogger = NewLoggerByWrapZap(
		DefaultAtomicLevel,
		Version,
		zapcore.AddSync(os.Stderr),
	)
}

func NewLoggerByWrapZap(level zapcore.LevelEnabler, version string, writes ...zapcore.WriteSyncer) *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendInt64(t.Unix())
	}

	var opts []zap.Option
	opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	if version != "" {
		opts = append(opts, zap.Fields(zap.String("v", version)))
	}

	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg),
			zap.CombineWriteSyncers(
				writes...,
			),
			level,
		),
		opts...,
	)
}

func Print(objs ...interface{}) {
	zapFields := make([]zap.Field, 0, len(objs))

	for i := 0; i < len(objs); i++ {
		zapFields = append(zapFields, zap.String(fmt.Sprintf("k%d", i), fmt.Sprintf("%#v", objs[i])))
	}
	DefaultLogger.Debug("", zapFields...)
}

func Printf(format string, a ...interface{}) {
	DefaultLogger.Debug("", zap.String("k0", fmt.Sprintf(format, a...)))
}

func CheckError(err error) {
	if err != nil {
		DefaultLogger.Panic(err.Error())
	}
}

func Debug(msg string, fields ...zap.Field) {
	DefaultLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	DefaultLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	DefaultLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	DefaultLogger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	DefaultLogger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	DefaultLogger.Fatal(msg, fields...)
}

func Sync() {
	DefaultLogger.Sync()
}
