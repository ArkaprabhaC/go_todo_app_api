package config

import (
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/todo_notes/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}
type DatabaseConfig struct {
	HostName string `mapstructure:"host_name"`
	Port     string `mapstructure:"port"`
	DbName   string `mapstructure:"db_name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	SslMode  string `mapstructure:"ssl_mode"`
}

const (
	configPath = "./config"
	configName = "config"
	configType = "json"
)

func ReadConfig() Config {
	log := logger.Logger()
	config := Config{}
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to read config file. %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config file. %s", err)
	}
	return config
}
