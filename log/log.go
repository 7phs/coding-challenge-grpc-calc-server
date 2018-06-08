package log

import (
		"log"
)

var (
	logLevel LogLevel = ALL
)

// Set logger level, but only before start a server to prevent sync problems.
func SetLogLevel(level LogLevel) {
	logLevel = level
}

// Get the current logger level.
func GetLogLevel() LogLevel {
	return logLevel
}

// Set prefix for the logger, but only before start a server to prevent sync problems.
func SetPrefix(p string) {
	log.SetPrefix(p)
}

// Log debugging messages.
func Debug(msgs ... interface{}) {
	if logLevel&DEBUG==0 {
		return
	}

	log.Println(msgs...)
}

// Log general messages.
func Info(msgs ... interface{}) {
	if logLevel&INFO==0 {
		return
	}

	log.Println(msgs...)
}

// Log warnings.
func Warning(msgs ... interface{}) {
	if logLevel&WARNING==0 {
		return
	}

	log.Println(msgs...)
}

// Log critical errors.
func Error(msgs ... interface{}) {
	if logLevel&ERROR==0 {
		return
	}

	log.Println(msgs...)
}

