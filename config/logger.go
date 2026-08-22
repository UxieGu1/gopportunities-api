package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Logger struct {
	service string
	level   string
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	level := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if level == "" {
		level = "info"
	}

	return &Logger{
		service: p,
		level:   level,
		writer:  os.Stdout,
	}
}

func (l *Logger) enabled(level string) bool {
	levels := map[string]int{"debug": 10, "info": 20, "warn": 30, "error": 40}
	current, ok := levels[l.level]
	if !ok {
		current = levels["info"]
	}
	wanted, ok := levels[level]
	if !ok {
		wanted = levels["info"]
	}
	return current <= wanted
}

func (l *Logger) write(level, message string, fields map[string]string) {
	if !l.enabled(level) {
		return
	}

	payload := map[string]interface{}{
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"level":     level,
		"service":   l.service,
		"message":   message,
	}
	for k, v := range fields {
		payload[k] = v
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprintf(l.writer, "{\"level\":\"%s\",\"service\":\"%s\",\"message\":\"%s\"}\n", level, l.service, message)
		return
	}

	fmt.Fprintln(l.writer, string(encoded))
}

func (l *Logger) Debug(v ...interface{}) {
	l.write("debug", fmt.Sprint(v...), nil)
}
func (l *Logger) Info(v ...interface{}) {
	l.write("info", fmt.Sprint(v...), nil)
}
func (l *Logger) Warn(v ...interface{}) {
	l.write("warn", fmt.Sprint(v...), nil)
}
func (l *Logger) Error(v ...interface{}) {
	l.write("error", fmt.Sprint(v...), nil)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.write("debug", fmt.Sprintf(format, v...), nil)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.write("info", fmt.Sprintf(format, v...), nil)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.write("warn", fmt.Sprintf(format, v...), nil)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.write("error", fmt.Sprintf(format, v...), nil)
}

func (l *Logger) InfoWithFields(message string, fields map[string]string) {
	l.write("info", message, fields)
}

func (l *Logger) ErrorWithFields(message string, fields map[string]string) {
	l.write("error", message, fields)
}