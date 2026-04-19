package logger

import (
	"io"
	"log"
	"os"
)

// ANSI color codes
const (
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorReset  = "\033[0m"
)

var (
	// Публичные логгеры
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
	Debug *log.Logger
)

func init() {
	// Создание логгеров с цветами
	Info = log.New(io.MultiWriter(os.Stdout), "INFO: ", log.Ldate|log.Ltime)
	Warn = log.New(io.MultiWriter(os.Stderr), colorYellow+"WARN: "+colorReset, log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr), colorRed+"ERROR: "+colorReset, log.Ldate|log.Ltime|log.Llongfile)
	Debug = log.New(io.MultiWriter(os.Stdout), colorBlue+"DEBUG: "+colorReset, log.Lmicroseconds|log.Lshortfile)
}
