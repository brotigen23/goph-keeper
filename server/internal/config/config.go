package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
)

func LoadDotEnv() error {
	return godotenv.Load()
}

type Config struct {
	Server struct {
		Address string `yaml:"Adress" env:"SERVER_ADDRESS" env-default:"http://localhost:8080/"`
	} `yaml:"server"`

	Database struct {
		User     string `yaml:"user" env:"DB_LOGIN"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		DBName   string `yaml:"dbname" env:"DB_NAME"`
	} `yaml:"database"`
}

func (c *Config) Load() error {
	err := cleanenv.ReadConfig("config/config.yml", c)
	if err != nil {
		return err
	}

	err = cleanenv.ReadEnv(c)
	if err != nil {
		return err
	}

	_, err = flags.Parse(c)
	if err != nil {
		return err
	}

	return nil
}

func (c Config) GetPostgresDSN() string {
	return fmt.Sprintf(
		`host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable`,
		c.Database.User, c.Database.Password, c.Database.DBName,
	)
}
