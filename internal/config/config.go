package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Http     Http     `yaml:"http"`
	Postgres Postgres `yaml:"postgres"`
}

type Http struct {
	Port int `yaml:"port"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	DBName   string `yaml:"db-name"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"ssl-mode"`
}

func NewConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Для выдачи API semantic version и других в будущем

func InitViperConfig() {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	// Enable Viper to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file, %v", err)
	}

}
