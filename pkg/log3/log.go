package log2

import (
	"context"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/go-spring/log"
	"github.com/pangu-2/go-tools/tools/datetimePg"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Level = slog.Level

const (
	LevelDebug = slog.LevelDebug
	LevelTrace = slog.Level(-2) // 自定义日志级别
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

type Logger struct {
	l   *slog.Logger
	lvl *slog.LevelVar // 用来动态调整日志级别
}

func NewDefault(level slog.Level) *Logger {
	return New(level, false)
}

func New(level slog.Level, jsonIs bool) *Logger {
	var lvl slog.LevelVar
	lvl.Set(level)
	log.Debugf(context.Background(), log.TagAppDef, "NewLogger.level:%+v ", level)
	date := datetimePg.Date()
	// 配置lumberjack用于按天分割日志
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./data/logs/" + date + ".log", // 日志文件路径
		MaxSize:    100,                            // 单个文件最大大小(MB)
		MaxBackups: 30,                             // 保留旧文件的最大数量
		MaxAge:     60,                             // 保留旧文件的最大天数
		Compress:   true,                           // 是否压缩/归档旧文件
		LocalTime:  true,                           // 使用本地时间而非UTC
	}
	// 创建MultiWriter，同时写入文件和控制台
	multiWriter := io.MultiWriter(lumberjackLogger, os.Stdout)
	var handler slog.Handler
	if jsonIs {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,

			// Level:     level, // 静态设置日志级别
			Level: &lvl, // 支持动态设置日志级别

			// 修改日志中的 Attr 键值对（即日志记录中附加的 key/value）
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.LevelKey {
					level := a.Value.Any().(slog.Level)
					levelLabel := level.String()

					switch level {
					case LevelTrace:
						// NOTE: 如果不设置，默认日志级别打印为 "level":"DEBUG+2"
						levelLabel = "TRACE"
					}

					a.Value = slog.StringValue(levelLabel)
				}

				// NOTE: 可以在这里修改时间输出格式
				// if a.Key == slog.TimeKey {
				//     if t, ok := a.Value.Any().(time.Time); ok {
				//         a.Value = slog.StringValue(t.Format(time.DateTime))
				//     }
				// }

				return a
			},
		})
	} else {
		opt := &slog.HandlerOptions{
			AddSource: true,
			Level:     &lvl, // 支持动态设置日志级别
		}
		handler = slog.NewTextHandler(multiWriter, opt)
		//handler = slog.NewTextHandler(os.Stdout, opt)
		// 创建自定义的文本处理器
		//handler = NewCustomTextHandler(os.Stdout, opt, "app")
	}
	h := slog.New(handler)
	slog.SetDefault(h)
	return &Logger{l: h, lvl: &lvl}
}

// SetLevel 动态调整日志级别
func (l *Logger) SetLevel(level Level) {
	l.lvl.Set(level)
}

func (l *Logger) Debug(msg string, args ...any) {
	// 不会走 *customlog.Logger.log() 调用，会走 *slog.Logger.log() 调用
	l.l.Debug(msg, args...)
}
func (l *Logger) Debugf(msg string, args ...any) {
	// 不会走 *customlog.Logger.log() 调用，会走 *slog.Logger.log() 调用
	l.l.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.Log(context.Background(), LevelInfo, msg, args...)
}
func (l *Logger) Infof(msg string, args ...any) {
	l.Log(context.Background(), LevelInfo, msg, args...)
}

// Trace 自定义的日志级别
func (l *Logger) Trace(msg string, args ...any) {
	l.Log(context.Background(), LevelTrace, msg, args...)
}

// Trace 自定义的日志级别
func (l *Logger) Tracef(msg string, args ...any) {
	l.Log(context.Background(), LevelTrace, msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.Log(context.Background(), LevelWarn, msg, args...)
}
func (l *Logger) Warnf(msg string, args ...any) {
	l.Log(context.Background(), LevelWarn, msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.Log(context.Background(), LevelError, msg, args...)
}
func (l *Logger) Errorf(msg string, args ...any) {
	l.Log(context.Background(), LevelError, msg, args...)
}

func (l *Logger) Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	l.log(ctx, level, msg, args...)
}

// log is the low-level logging method for methods that take ...any.
// It must always be called directly by an exported logging method
// or function, because it uses a fixed call depth to obtain the pc.
func (l *Logger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if !l.l.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	// NOTE: 这里修改 skip 为 4，*slog.Logger.log 源码中 skip 为 3
	runtime.Callers(4, pcs[:])
	pc = pcs[0]
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = l.l.Handler().Handle(ctx, r)
}
