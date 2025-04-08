package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
)

func LoadDotEnv() error {
	err := godotenv.Load()
	return err
}

type Config struct {
	Server struct {
		Address string `yaml:"address" env:"SERVER_ADDRESS" env-default:"http://localhost:8080/"`
	} `yaml:"server"`

	Database struct {
		User     string `yaml:"user" env:"DB_LOGIN"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		DBName   string `yaml:"dbname" env:"DB_NAME"`
	} `yaml:"database"`

	JWT struct {
		AccessKey  string `yaml:"access_key"`
		RefreshKey string `yaml:"refresh_key"`
	} `yaml:"JWT"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) Default() (*Config, error) {
	c.Load()
	return c, nil
}

func (c *Config) Load() error {
	err := LoadDotEnv()
	if err != nil {
		log.Println(err)
	}
	err = cleanenv.ReadConfig("config/config.yml", c)
	if err != nil {
		log.Println(err)
	}

	err = cleanenv.ReadEnv(c)
	if err != nil {
		log.Println(err)
	}

	_, err = flags.Parse(c)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (c Config) GetPostgresDSN() string {
	return fmt.Sprintf(
		`host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable`,
		c.Database.User, c.Database.Password, c.Database.DBName,
	)
}
