package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const Path = ".gob.yaml"

type Config struct {
	Database   Database   `yaml:"database"`
	Migrations Migrations `yaml:"migrations"`
	isCli      bool       `yaml:"-"`
}

type Database struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func (cfg *Database) URL() (string, error) {
	switch cfg.Driver {
	case "postgres":
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Name,
		), nil
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Name,
		), nil
	default:
		return "", fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}
}

type Migrations struct {
	Dir string `yaml:"dir"`
	Ext string `yaml:"ext"`
	Seq bool   `yaml:"seq"`
}

func Load() (Config, error) {
	return LoadByPath(Path)
}

func LoadByPath(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}
