package logger

import (
	"fmt"
	"io"
	"log"

	"github.com/hashicorp/logutils"
)

func Initialise() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func Configure(output io.Writer, debugEnabled bool) {
	Initialise()
	minLevel := "INFO"
	if debugEnabled {
		minLevel = "DEBUG"
	}
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		Writer:   output,
		MinLevel: logutils.LogLevel(minLevel),
	}
	log.SetOutput(filter)

}

func Trace(message string, args ...interface{}) {
	logLine("TRACE", fmt.Sprintf(message, args...))
}

func Debug(message string, args ...interface{}) {
	logLine("DEBUG", fmt.Sprintf(message, args...))
}

func Info(message string, args ...interface{}) {
	logLine("INFO", fmt.Sprintf(message, args...))
}

func Warn(message string, args ...interface{}) {
	logLine("WARN", fmt.Sprintf(message, args...))
}

func Error(message string, args ...interface{}) {
	logLine("ERROR", fmt.Sprintf(message, args...))
}

func Fatal(message string, args ...interface{}) {
	logLine("FATAL", fmt.Sprintf(message, args...))
}

func Print(message string, args ...interface{}) {
	_, _ = fmt.Fprintf(log.Writer(), message, args...)
}

func logLine(level, message string) {
	log.Printf("[%s] %s", level, message)
}
