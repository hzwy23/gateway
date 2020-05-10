package config

import "time"

type RegisterCenter struct {
	 RefreshFrequency time.Duration `yaml:"refreshFrequency"`
	 EurekaConfig struct {
		ServiceUrls []string `yaml:"serviceUrls"`
	}
}
