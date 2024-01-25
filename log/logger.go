package logging

import (
	"log"
	"os"
)

type Logger struct {
	Info *log.Logger
	Warn *log.Logger
	Error *log.Logger
}

func NewLogger(f int) *Logger{
	l := &Logger{}
	l.Info = log.New(os.Stdout, "[INFO] ", f)
	l.Warn = log.New(os.Stdout, "[WARN] ", f)
	l.Error = log.New(os.Stdout, "[ERROR] ", f)

	return l
}