package config

import "time"

type ServerConfig struct {
	Port        int           `yaml:"port"`
	ContextPath string        `yaml:"contextPath"`
	Timeout     time.Duration `yaml:"timeout"`
	Host        string        `yaml:"host"`
}
