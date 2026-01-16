package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type ContextConfig struct {
	TimeOut time.Duration `yaml:"timeout"`
}

type ServerConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	IdleTimeout     time.Duration `yaml:"idle_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	Context         ContextConfig `yaml:"context"`
}

type PostgresPoolConfig struct {
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
}
type PostgresConfig struct {
	Host     string             `yaml:"host"`
	User     string             `yaml:"user" `
	Password string             `yaml:"password"`
	DB       string             `yaml:"db"`
	SSLMode  string             `yaml:"sslmode"`
	TimeZone string             `yaml:"time_zone"`
	Pool     PostgresPoolConfig `yaml:"pool"`
	Port     int                `yaml:"port" `
}

type Config struct {
	Env      string         `yaml:"env"`
	LogLevel string         `yaml:"log_level"`
	Postgres PostgresConfig `yaml:"postgres"`
	Server   ServerConfig   `yaml:"server"`
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
