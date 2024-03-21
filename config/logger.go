package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	write   io.Writer
}

func NewLogger(prefix string) *Logger {
	write := io.Writer(os.Stdout)
	logger := log.New(write, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(write, "DEBUG: ", logger.Flags()),
		info:    log.New(write, "INFO: ", logger.Flags()),
		warning: log.New(write, "WARNING: ", logger.Flags()),
		err:     log.New(write, "ERROR: ", logger.Flags()),
		write:   write,
	}
}

func (logger *Logger) Debug(values ...interface{}) {
	logger.debug.Println(values...)
}

func (logger *Logger) Info(values ...interface{}) {
	logger.info.Println(values...)
}

func (logger *Logger) Warning(values ...interface{}) {
	logger.warning.Println(values...)
}

func (logger *Logger) Error(values ...interface{}) {
	logger.err.Println(values...)
}

func (logger *Logger) DebugFormat(format string, values ...interface{}) {
	logger.debug.Printf(format, values...)
}

func (logger *Logger) InfoFormat(format string, values ...interface{}) {
	logger.info.Printf(format, values...)
}

func (logger *Logger) WarningFormat(format string, values ...interface{}) {
	logger.warning.Printf(format, values...)
}

func (logger *Logger) ErrorFormat(format string, values ...interface{}) {
	logger.err.Printf(format, values...)
}
