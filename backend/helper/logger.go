package helper

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// zap - カスタム現地時間フォーマット
func ZapLogLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// zap logger for gorm
type GormZapLogger struct {
	log *zap.Logger
	logger.Config
	normalStr, traceStr, traceErrStr, traceWarnStr string
}

var Log *GormZapLogger
var RuntimeRoot string

func InitLogger() {
	now := time.Now()
	logPath := "../logs"
	logFileName := fmt.Sprintf("%s/%04d-%02d-%02d.log", logPath, now.Year(), now.Month(), now.Day())
	hook := &lumberjack.Logger{
		Filename:   logFileName, // ログファイルのパス
		MaxSize:    50,          // ファイルサイズ, M
		MaxBackups: 100,         // ログファイルの最大数
		MaxAge:     30,          // ログの保持時間、日数
		Compress:   false,       // 圧縮するかどうか
	}
	defer func(hook *lumberjack.Logger) {
		err := hook.Close()
		if err != nil {
			fmt.Printf("[InitLogger] func(hook *lumberjack.Logger) クローズに失敗しました")
		}
	}(hook)

	enConfig := zap.NewProductionEncoderConfig()
	enConfig.EncodeTime = ZapLogLocalTimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(enConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)),
		//Print Log to the console
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),),
		// log level(-1:Debug, 0:Info, -1<=level<=5, refer to zap.level)
		zapcore.InfoLevel, // 设置日志级别为 Info
	)

	l := zap.New(core)
	Log = NewGormZapLogger(l, logger.Config{})
	Log.Debug(context.Background(), "初期化ログ機能が完了しました。")
}

// New logger like gorm2  生成logs对象
func NewGormZapLogger(zapLogger *zap.Logger, config logger.Config) *GormZapLogger {

	_, file, _, _ := runtime.Caller(0)
	RuntimeRoot = strings.TrimSuffix(file, "main.go")

	var (
		normalStr    = "%v "
		traceStr     = "%v\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%v %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%v %s\n[%.3fms] [rows:%v] %s"
	)
	l := &GormZapLogger{
		log:          zapLogger,
		Config:       config,
		normalStr:    normalStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
	return l
}

func (l *GormZapLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func removePrefix(s string) string {
	s = strings.TrimPrefix(s, RuntimeRoot)
	return s
}

// Debug print info
func (l GormZapLogger) Debug(ctx context.Context, format string, args ...interface{}) {
	if l.log.Core().Enabled(zapcore.DebugLevel) {
		l.log.Sugar().Debugf(l.normalStr+format, append([]interface{}{removePrefix(utils.FileWithLineNum())}, args...)...)
	}
}

func (l GormZapLogger) Info(ctx context.Context, format string, args ...interface{}) {
	if l.log.Core().Enabled(zapcore.InfoLevel) {
		l.log.Sugar().Infof(l.normalStr+format, append([]interface{}{removePrefix(utils.FileWithLineNum())}, args...)...)
	}
}

func (l GormZapLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	if l.log.Core().Enabled(zapcore.WarnLevel) {
		l.log.Sugar().Warnf(l.normalStr+format, append([]interface{}{removePrefix(utils.FileWithLineNum())}, args...)...)
	}
}

func (l GormZapLogger) Error(ctx context.Context, format string, args ...interface{}) {

	if l.log.Core().Enabled(zapcore.ErrorLevel) {
		l.log.Sugar().Errorf(l.normalStr+format, append([]interface{}{removePrefix(utils.FileWithLineNum())}, args...)...)
	}
}

// Trace print sql message
func (l GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if !l.log.Core().Enabled(zapcore.DPanicLevel) || l.LogLevel <= logger.Silent {
		return
	}
	lineNum := removePrefix(utils.FileWithLineNum())
	elapsed := time.Since(begin)
	elapsedF := float64(elapsed.Nanoseconds()) / 1e6
	sql, rows := fc()
	row := "-"
	if rows > -1 {
		row = fmt.Sprintf("%d", rows)
	}
	switch {
	case l.log.Core().Enabled(zapcore.ErrorLevel) && err != nil && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		l.log.Error(fmt.Sprintf(l.traceErrStr, lineNum, err, elapsedF, row, sql))
	case l.log.Core().Enabled(zapcore.WarnLevel) && elapsed > l.SlowThreshold && l.SlowThreshold != 0:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		l.log.Warn(fmt.Sprintf(l.traceWarnStr, lineNum, slowLog, elapsedF, row, sql))
	case l.log.Core().Enabled(zapcore.DebugLevel):
		l.log.Debug(fmt.Sprintf(l.traceStr, lineNum, elapsedF, row, sql))
	case l.log.Core().Enabled(zapcore.InfoLevel):
		l.log.Info(fmt.Sprintf(l.traceStr, lineNum, elapsedF, row, sql))
	}
}
