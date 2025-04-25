package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env" env:"ENV" env-default:"local"`
	HTTPServer HTTPServer `yaml:"http_server"`
	DB         DBConfig   `yaml:"db"`
}
type DBConfig struct {
	Host string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	User string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Pass string `yaml:"pass" env:"DB_PASS" env-default:"1234"`
	Name string `yaml:"name" env:"DB_NAME" env-default:"postgres"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" env:"HOST" end-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"4"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60"` // Idle timeout in seconds
}

func MustLoad() Config {
	// Load environment variables or from .yaml file
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config/local.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %s", configPath)
	} // error handling

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to read config file: %s", err)
	}

	return cfg
}
