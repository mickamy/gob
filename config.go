package gob

const ConfigPath = ".gob.yaml"

type Config struct {
	Database   Database   `yaml:"database"`
	Migrations Migrations `yaml:"migrations"`
}

type Database struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Migrations struct {
	Dir string `yaml:"dir"`
}
