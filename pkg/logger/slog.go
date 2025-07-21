package logger

import (
	"log/slog"
	"os"
)

type SlogLogger struct{
	log *slog.Logger
}

const (
	envProd = "prod"
	envDebug = "debug"
)

func NewSlogLogger(level string) *SlogLogger{
	var log *slog.Logger

	switch level{
	case envDebug:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level : slog.LevelDebug}),)
	case envProd:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level : slog.LevelInfo}),)
	}

	return &SlogLogger{log : log}

}

func (l *SlogLogger) Info(msg string, fields ...interface{}){
	l.log.Info(msg, fields...)
}

func (l *SlogLogger) Debug(msg string, fields ...interface{}){
	l.log.Debug(msg, fields...)
}

func (l *SlogLogger) Error(msg string, fields ...interface{}){
	l.log.Error(msg, fields...)
}

func (l *SlogLogger) Warn(msg string, fields ...interface{}){
	l.log.Warn(msg, fields...)
}