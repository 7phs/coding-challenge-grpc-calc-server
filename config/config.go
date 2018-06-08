package config

import (
	"os"

	"github.com/7phs/coding-challenge-grpc-calc-server/log"
	"strconv"
	"fmt"
	"time"
)

const (
	DEFAULT_PORT      = 9090
	DEFAULT_TIMEOUT   = 1 * time.Second
	DEFAULT_LOG_LEVEL = log.INFO

	CONFIG_PORT      = "PORT"
	CONFIG_TIMEOUT   = "TIMEOUT"
	CONFIG_LOG_LEVEL = "LOG_LEVEL"
)

// Config is a struct storing both side configuration parameters - a service and a client.
type Config struct {
	port     int
	timeout  time.Duration
	logLevel log.LogLevel
}

// Return a calc service port number.
func (o *Config) Port() int {
	return o.port
}

// Set a calc service port number. Useful in testing.
func (o *Config) SetPort(port int) *Config {
	o.port = port

	return o
}

// Set a request timeout. Useful in testing.
func (o *Config) SetTimeout(timeout time.Duration) *Config {
	o.timeout = timeout

	return o
}

// Return a port formatted as an address string.
func (o *Config) Address() string {
	return fmt.Sprintf(":%d", o.port)
}

// Return a log level.
func (o *Config) LogLevel() log.LogLevel {
	return o.logLevel
}

// Return a request timeout.
func (o *Config) Timeout() time.Duration {
	return o.timeout
}

// Parsing environment variables and init a configuration.
func ParseConfig() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv(CONFIG_PORT))
	if err != nil || port == 0 {
		port = DEFAULT_PORT
	}

	timeout := DEFAULT_TIMEOUT
	millis, err := strconv.Atoi(os.Getenv(CONFIG_TIMEOUT))
	if err == nil && millis > 0 {
		timeout = time.Duration(millis) * time.Millisecond
	}

	logLevel := log.ParseLevel(os.Getenv(CONFIG_LOG_LEVEL), DEFAULT_LOG_LEVEL)

	return &Config{
		port:     int(port),
		timeout:  timeout,
		logLevel: logLevel,
	}, nil
}
