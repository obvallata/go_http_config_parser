package parse_http

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	Port    int           `yaml:"port"`
	Host    string        `yaml:"host"`
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

const (
	DefaultTimeout = 5 * time.Second
)

func ParseHTTPConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("can't read file %s due to %s", filename, err)
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Printf("read config file error: %s", err)
		return nil, err
	}
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}
	return &config, nil
}
