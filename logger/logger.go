package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/biisal/fast-stream-bot/config"
)

type ShortSourceHandler struct {
	slog.Handler
}

type HandlerType string

const (
	FileHandler   HandlerType = "fileHandler"
	StdoutHandler HandlerType = "stdoutHandler"
	HandlerMulti  HandlerType = "multiHandler"
)

func (h ShortSourceHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.PC != 0 {
		frames := runtime.CallersFrames([]uintptr{r.PC})
		frame, _ := frames.Next()
		if frame.File != "" {
			filename := filepath.Base(frame.File)
			r.AddAttrs(slog.String("source", filename+":"+strconv.Itoa(frame.Line)))
		}
	}
	return h.Handler.Handle(ctx, r)
}

// Required methods for a proper slog.Handler
func (h ShortSourceHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h ShortSourceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return ShortSourceHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h ShortSourceHandler) WithGroup(name string) slog.Handler {
	return ShortSourceHandler{Handler: h.Handler.WithGroup(name)}
}

// MultiHandler writes to multiple handlers
type MultiHandler struct {
	handlers []slog.Handler
}

func (m MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, h := range m.handlers {
		_ = h.Handle(ctx, r) // fire and forget (or collect errors if you want)
	}
	return nil
}

func (m MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	copies := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		copies[i] = h.WithAttrs(attrs)
	}
	return MultiHandler{handlers: copies}
}

func (m MultiHandler) WithGroup(name string) slog.Handler {
	copies := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		copies[i] = h.WithGroup(name)
	}
	return MultiHandler{handlers: copies}
}

func SetupSlog(env string, handlerType HandlerType) (io.Closer, error) {
	var level slog.Leveler = slog.LevelError
	if env == config.ENVIRONMENT_LOCAL {
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: false,
	}

	var logFile *os.File
	var handler slog.Handler

	switch handlerType {
	case FileHandler, HandlerMulti:
		f, err := os.OpenFile("fsb.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		logFile = f
	}

	switch handlerType {
	case FileHandler:
		handler = slog.NewTextHandler(logFile, opts)

	case StdoutHandler:
		handler = slog.NewTextHandler(os.Stdout, opts)

	case HandlerMulti:
		fileH := slog.NewTextHandler(logFile, opts)
		stdoutH := slog.NewTextHandler(os.Stdout, opts)
		handler = MultiHandler{handlers: []slog.Handler{fileH, stdoutH}}
	}

	wrapped := ShortSourceHandler{Handler: handler}
	slog.SetDefault(slog.New(wrapped))

	return logFile, nil
}
