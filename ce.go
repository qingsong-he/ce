package ce

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// code commit hash
var DefaultVersion string

// from that running module
var DefaultFrom string

var DefaultLogger *zap.Logger
var DefaultAtomicLevel zap.AtomicLevel

func init() {
	DefaultAtomicLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	DefaultLogger = NewLoggerByWrapZap(
		DefaultAtomicLevel,
		zap.PanicLevel,
		DefaultFrom,
		DefaultVersion,
		zapcore.AddSync(os.Stderr),
	)
}

func NewLoggerByWrapZap(level, levelByStacktrace zapcore.LevelEnabler, from, version string, writes ...zapcore.WriteSyncer) *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	var opts []zap.Option
	opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(levelByStacktrace))
	if from != "" {
		opts = append(opts, zap.Fields(zap.String("f", from)))
	}
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

type panicByMe struct {
	OriginalErr error
}

func (p *panicByMe) Error() string {
	return p.OriginalErr.Error()
}

func IsFromMe(errByPanic interface{}) (*panicByMe, bool) {
	me, ok := errByPanic.(*panicByMe)
	return me, ok
}

func CheckError(err error) {
	if err != nil {
		DefaultLogger.Error("CheckError", zap.Error(err), zap.Stack("callStack"))
		panic(&panicByMe{OriginalErr: err})
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
