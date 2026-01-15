package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Host string `yaml:"host" env:"HOST" env-required:"true"`
	Port int    `yaml:"port" env:"PORT" env-required:"true"`
}

type PostgresPoolConfig struct {
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
}
type PostgresConfig struct {
	Host     string             `yaml:"host"`
	Port     int                `yaml:"port" `
	User     string             `yaml:"user" `
	Password string             `yaml:"password"`
	DB       string             `yaml:"db"`
	SSLMode  string             `yaml:"sslmode"`
	TimeZone string             `yaml:"time_zone"`
	Pool     PostgresPoolConfig `yaml:"pool"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server" env-prefix:"SERVER_"`
	Postgres PostgresConfig `yaml:"postgres" env-prefix:"POSTGRES_"`
}

func New(configPath string) (*Config, error) {
	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
