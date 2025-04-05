package logger

import (
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Logger
type Logger struct {
	logger *slog.Logger

	logFile *os.File
}

func New() *Logger {
	return &Logger{
		logger: nil,
	}
}

func (l *Logger) Default() *Logger {
	logPath := filepath.Join(".", "log")
	logFileName := logPath + "/" + time.Now().String() + ".log"

	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		return nil
	}
	l.logFile, err = os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil
	}

	l.logger = slog.New(slog.NewJSONHandler(l.logFile, nil))

	return l
}

func (l Logger) Error(err error) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		l.logger.Error("Error occurred", slog.String("error", "could not retrieve file/line"))
		return
	}
	l.logger.Error(
		"Error occured",
		"error", err.Error(),
		slog.String("file", file), slog.Int("line", line))
}

func (l Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l Logger) CloseFile() {
	l.logFile.Close()
}
