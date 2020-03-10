package config

import "github.com/jinzhu/configor"

var Config struct {
	DB struct {
		Name		string `env:"DBName"`
		Host		string `env:"DBHost"`
		Port     int    `env:"DBPort"`
		User     string `env:"DBUser"`
		Password string `env:"DBPassword"`
	}
} 

func init() {
	if err := configor.Load(&Config, "config/config.yml"); err != nil {
		panic(err)
	}
}