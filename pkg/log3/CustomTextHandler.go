package log2

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strings"
	"sync"
)

// 自定义时间格式
const timeFormat = "2006-01-02 15:04:05.000"

// 自定义处理器结构体
type CustomTextHandler struct {
	opts    slog.HandlerOptions
	writer  io.Writer
	mu      sync.Mutex
	channel string
}

// 工厂函数创建自定义处理器
func NewCustomTextHandler(w io.Writer, opts *slog.HandlerOptions, channel string) *CustomTextHandler {
	h := &CustomTextHandler{
		writer:  w,
		channel: channel,
	}
	if opts != nil {
		h.opts = *opts
	}
	return h
}

// 实现 slog.Handler 接口的 Enabled 方法
func (h *CustomTextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	if h.opts.Level != nil {
		return level >= h.opts.Level.Level()
	}
	return true
}

// 实现 slog.Handler 接口的 Handle 方法
func (h *CustomTextHandler) Handle(ctx context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	// 准备日志元素
	var (
		datetime = r.Time.Format(timeFormat)
		level    = r.Level.String()
		caller   = ""
		message  = r.Message
		data     = make(map[string]any)
		extra    = ""
	)

	// 获取调用者信息
	if r.PC != 0 && h.opts.AddSource {
		if pc, file, line, ok := runtime.Caller(6); ok {
			f := runtime.FuncForPC(pc)
			if f != nil {
				caller = fmt.Sprintf("%s:%d", file, line)
			}
		}
	}

	// 处理所有属性
	r.Attrs(func(attr slog.Attr) bool {
		// 跳过已处理的标准字段
		if attr.Key == slog.TimeKey || attr.Key == slog.LevelKey || attr.Key == slog.MessageKey {
			return true
		}

		// 特殊处理 extra 字段
		if attr.Key == "extra" {
			extra = fmt.Sprintf("%v", attr.Value.Any())
			return true
		}

		// 其他字段放入 data
		data[attr.Key] = attr.Value.Any()
		return true
	})

	// 构建数据部分的字符串
	var dataStr strings.Builder
	for k, v := range data {
		if dataStr.Len() > 0 {
			dataStr.WriteString(" ")
		}
		dataStr.WriteString(fmt.Sprintf("%s=%v", k, v))
	}
	//模版 [{{datetime}}] [{{channel}}] [{{level}}] [{{caller}}] {{message}} {{data}} {{extra}}\n
	// 构建最终日志行
	logLine := fmt.Sprintf("[%s] [%s] [%s] [%s] %+v %+v %+v\n",
		datetime, h.channel, level, caller, message, dataStr.String(), extra)

	// 写入日志
	_, err := h.writer.Write([]byte(logLine))
	return err
}

// 实现 slog.Handler 接口的 WithAttrs 方法
func (h *CustomTextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// 创建新的处理器并复制属性
	newHandler := *h
	return &newHandler
}

// 实现 slog.Handler 接口的 WithGroup 方法
func (h *CustomTextHandler) WithGroup(name string) slog.Handler {
	// 创建新的处理器并设置组名
	newHandler := *h
	return &newHandler
}
