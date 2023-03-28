package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewLogger() *Logger {
	// Open a file for writing
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}

	// Create a new logger instance that writes to the file
	infoWriter := io.MultiWriter(os.Stdout, file)
	errorWriter := io.MultiWriter(os.Stderr, file)

	infoLogger := log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime)
	errorLogger := log.New(errorWriter, "ERROR: ", log.Ldate|log.Ltime)

	return &Logger{
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	filename := path.Base(file)
	l.infoLogger.Printf("File: %s, Line: %d - "+format, append([]interface{}{filename, line}, args...)...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	filename := path.Base(file)
	l.errorLogger.Printf("File: %s, Line: %d - "+format, append([]interface{}{filename, line}, args...)...)
}
