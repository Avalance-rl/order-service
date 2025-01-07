package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	GRPCServer GRPCServer `yaml:"grpc_server"`
	Database   Database   `yaml:"database"`
	Redis      Redis      `yaml:"redis"`
}

type GRPCServer struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"ssl_mode"`
	MaxConns int    `yaml:"max_conns"`
}

type Redis struct{}

func Load(configPath string) (Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, err
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
