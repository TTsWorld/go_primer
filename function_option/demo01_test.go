package function_option

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

type Config struct {
	cfg1 int
	cfg2 int
	cfg3 int
}

type Option func(*Config)

func WithCfg1(c int) Option {
	return func(cfg *Config) {
		cfg.cfg1 = c
	}
}

func WithCfg2(c int) Option {
	return func(cfg *Config) {
		cfg.cfg2 = c
	}
}

func WithCfg3(c int) Option {
	return func(cfg *Config) {
		cfg.cfg3 = c
	}
}

func NewConfig(cfg1, cfg2 int, options ...func(config *Config)) (*Config, error) {
	config := Config{
		cfg1: 1,
		cfg2: 2,
		cfg3: 3,
	}

	for _, option := range options {
		option(&config)
	}
	return &config, nil
}

func TestFuntionOption(t *testing.T) {
	cfg, err := NewConfig(1, 2, WithCfg3(3))
	if err != nil {
		fmt.Printf("new Config:%v", err)
		return
	}
	debugInfo, _ := jsoniter.MarshalToString(cfg)
	fmt.Printf("%+v", debugInfo)

}
