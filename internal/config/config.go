package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	General  GeneralConfig  `yaml:"general"`
	Metrics  MetricsConfig  `yaml:"metrics"`
	Database DatabaseConfig `yaml:"database"`
	Tracker  TrackerConfig  `yaml:"tracker"`
	Ticket   TicketConfig   `yaml:"ticket"`
}

type GeneralConfig struct {
	Port int `yaml:"port"`
}

type MetricsConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Path string `yaml:"path"`
}

type TrackerConfig struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

type TicketConfig struct {
	Project   ProjectConfig   `yaml:"project"`
	Issuetype IssuetypeConfig `yaml:"issuetype"`
	Labels    []string        `yaml:"labels"`
	Priority  PriorityConfig  `yaml:"priority"`
}

type ProjectConfig struct {
	Key string `yaml:"key"`
}

type IssuetypeConfig struct {
	Name string `yaml:"name"`
}

type PriorityConfig struct {
	Name string `yaml:"name"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &cfg, nil
}
