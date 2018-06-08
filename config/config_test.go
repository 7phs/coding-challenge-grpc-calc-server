package config

import (
	"os"
	"testing"
	"time"
	"github.com/7phs/coding-challenge-grpc-calc-server/log"
	"fmt"
)

func SetUpParseCofigParameter() func() {
	prev := map[string]string{}

	for _, name := range []string{CONFIG_PORT, CONFIG_TIMEOUT, CONFIG_LOG_LEVEL} {
		prev[name] = os.Getenv(name)
	}

	return func() {
		for name, value := range prev {
			os.Setenv(name, value)
		}
	}
}

func TestParseConfig(t *testing.T) {
	defer SetUpParseCofigParameter()()

	testSuites := []*struct {
		port             string
		expectedPort     int
		expectedAddress  string
		timeout          string
		expectedTimeout  time.Duration
		logLevel         string
		expectedLogLevel log.LogLevel
	}{
		{
			expectedPort: DEFAULT_PORT, expectedAddress: fmt.Sprint(":", DEFAULT_PORT), expectedTimeout: DEFAULT_TIMEOUT, expectedLogLevel: DEFAULT_LOG_LEVEL,
		},
		{
			port: "assaddsa", timeout: "-234234", logLevel: "fert",
			expectedPort: DEFAULT_PORT, expectedAddress: fmt.Sprint(":", DEFAULT_PORT), expectedTimeout: DEFAULT_TIMEOUT, expectedLogLevel: DEFAULT_LOG_LEVEL,
		},
		{
			port: "3040", timeout: "120", logLevel: "error",
			expectedPort: 3040, expectedAddress: ":3040", expectedTimeout: 120 * time.Millisecond, expectedLogLevel: log.ERROR,
		},
		//{
		//	client:      ":sdfsdf", expectedClient: DEFAULT_CLIENT,
		//	eventSource: ":dasdas", expectedEventSrc: DEFAULT_EVENT_SOURCE,
		//	queueLimit:  "dkjadslj", expectedQueueLimit: DEFAULT_QUEUE_LIMIT,
		//	queueTTL:    "asdasd", expectedQueueTTL: DEFAULT_QUEUE_TTL * time.Millisecond,
		//	logLevel:    "unksljkdf", expectedLogLevel: logger.CalcLevel(DEFAULT_LOG_LEVEL),
		//},
		//{
		//	client:      ":9090", expectedClient: ":9090",
		//	eventSource: ":7777", expectedEventSrc: ":7777",
		//	queueLimit:  "8888", expectedQueueLimit: 8888,
		//	queueTTL:    "700", expectedQueueTTL: 700 * time.Millisecond,
		//	logLevel:    "error", expectedLogLevel: logger.ERROR,
		//},
	}

	for i, test := range testSuites {
		os.Setenv(CONFIG_PORT, test.port)
		os.Setenv(CONFIG_TIMEOUT, test.timeout)
		os.Setenv(CONFIG_LOG_LEVEL, test.logLevel)

		params, err := ParseConfig()

		if err != nil {
			t.Error("failed to parse config")
		}

		if exist := params.Port(); exist != test.expectedPort {
			t.Error(i, ": failed to parse environment param for port. Got '", exist, "', but expected is '", test.expectedPort, "'")
		}

		if exist := params.Timeout(); exist != test.expectedTimeout {
			t.Error(i, ": failed to parse environment param for timeout. Got '", exist, "', but expected is '", test.expectedTimeout, "'")
		}

		if exist := params.Address(); exist != test.expectedAddress {
			t.Error(i, ": failed to parse environment param for address. Got '", exist, "', but expected is '", test.expectedTimeout, "'")
		}

		if exist := params.LogLevel(); exist != log.CalcLevel(test.expectedLogLevel) {
			t.Error(i, ": failed to parse environment param for log level. Got '", exist, "', but expected is '", test.expectedLogLevel, "'")
		}
	}
}
