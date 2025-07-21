package config

import (
	"github.com/ilyakaznacheev/cleanenv"

	"log"
	"os"

)

type Config struct{
	Env string `yaml:"env" env-default: "prod"`
	Repository `yaml:"repository"`
	HTTPServer `yaml:"httpserver"`
}

type Repository struct{
	Host string `yaml: "host" env-default: "localhost"`
	Port string `yaml: "port" env-default: "5432"`
	User string `yaml: "user"`
	Password string `yaml: "password"`
	Name string `yaml: "name"`
}

type HTTPServer struct{
	Address string `yaml: "address" env-default: "localhost:8090"`
}

func NewConfig() *Config{
	config_path := "/home/thethosaighfisic/marketplace/cmd/app/config/config.yaml"
	if config_path == ""{
		log.Fatal("CONFIG_PATH is empty")
	}

	_, err := os.Stat(config_path)
	if os.IsNotExist(err) {
		log.Fatalf("config_path is not exist: %s", config_path)
	}

	var cfg Config
	err = cleanenv.ReadConfig(config_path, &cfg)
	if err != nil {
		log.Fatalf("read config failed: %s", err)
	}
	return &cfg
}