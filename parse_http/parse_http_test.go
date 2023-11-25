package parse_http

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type TestConfigResult struct {
	name     string
	filename string
	expected *Config
	errOccur bool
}

func TestParseHTTPConfig(t *testing.T) {
	assertLocal := assert.New(t)
	tests := []TestConfigResult{
		{name: "SimpleConfig",
			filename: "./src/simple_config.yaml",
			expected: &Config{Port: 1234, Host: "127.0.0.2", Timeout: 30 * time.Second}},
		{name: "NoTimeoutConfig",
			filename: "./src/no_timeout_config.yaml",
			expected: &Config{Port: 1234, Host: "127.0.0.2", Timeout: DefaultTimeout}},
		{name: "NoConfig",
			filename: "./src/non_existing.yaml",
			expected: nil, errOccur: true},
		{name: "BadConfig",
			filename: "./src/bad_config.yaml",
			expected: nil, errOccur: true},
	}
	for _, test := range tests {
		got, err := ParseHTTPConfig(test.filename)
		assertLocal.Equal(test.expected, got, test.name)
		if test.errOccur {
			assertLocal.NotEqual(err, nil)
		}
	}
}
