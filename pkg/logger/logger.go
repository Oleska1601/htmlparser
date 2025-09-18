package logger

import (
	"log/slog"
	"os"
	"strings"
)

type Logger struct {
	logger *slog.Logger
}

func New(level string) *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	var lev slog.Level
	switch strings.ToLower(level) {
	case "debug":
		lev = slog.LevelDebug
	case "info":
		lev = slog.LevelInfo
	case "warn":
		lev = slog.LevelWarn
	case "error":
		lev = slog.LevelError
	default:
		lev = slog.LevelInfo
	}
	slog.SetLogLoggerLevel(lev)
	slog.SetDefault(logger)
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
