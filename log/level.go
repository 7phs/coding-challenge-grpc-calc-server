package log

import "strings"

type LogLevel int

const (
	DEBUG   LogLevel = 0x1
	INFO    LogLevel = 0x2
	WARNING LogLevel = 0x4
	ERROR   LogLevel = 0x8

	ALL LogLevel = DEBUG | INFO | WARNING | ERROR
)

// Parse a string as log level name to numeric representation.
func ParseLevel(level string, edgeLevel LogLevel) LogLevel {
	switch strings.ToLower(level) {
	case "debug":
		edgeLevel = DEBUG
	case "info":
		edgeLevel = INFO
	case "warning":
		edgeLevel = WARNING
	case "error":
		edgeLevel = ERROR
	}

	return CalcLevel(edgeLevel)
}

// Calculates a mask of logger from an minimum edge log level.
func CalcLevel(edgeLevel LogLevel) (result LogLevel) {
	start := false

	for _, v := range []LogLevel{DEBUG, INFO, WARNING, ERROR} {
		if v == edgeLevel {
			start = true
		}

		if start {
			result |= v
		}
	}

	return result
}

// Getting a minimum edge numeric representation of a log level.
func EdgeLevel(level LogLevel) LogLevel {
	for _, v := range []LogLevel{DEBUG, INFO, WARNING, ERROR} {
		if level&v!=0 {
			return v
		}
	}

	return 0
}

// String representation of a log level.
func (level LogLevel) String() string {
	switch EdgeLevel(level) {
	case DEBUG:
		return "Debug"
	case INFO:
		return "Info"
	case WARNING:
		return "Warning"
	case ERROR:
		return "Error"
	}

	return "Unknown"
}
