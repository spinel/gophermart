package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

const (
	defaultRunAddress       = "localhost:8080"
	defaultDatabaseURI      = "postgres://postgres:postgres@localhost:5439/postgres?sslmode=disable"
	defaultPgMigrationsPath = "file://internal/app/repository/pg/migrations"
)

// Config of app
type Config struct {
	JSON *viper.Viper

	HTTPAddr         string `envconfig:"RUN_ADDRESS"`
	LogLevel         string `envconfig:"LOG_LEVEL"`
	PgURL            string `envconfig:"DATABASE_URI"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		if err := initViper(); err != nil {
			log.Fatal(err)
		}

		setDefault(&config)

		fmt.Println("Configuration:", string(configBytes))
		config.JSON = viper.GetViper()
	})
	return &config
}

func setDefault(c *Config) {
	if c.HTTPAddr == "" {
		c.HTTPAddr = defaultRunAddress
	}
	if c.PgURL == "" {
		c.PgURL = defaultDatabaseURI
	}
	if c.PgMigrationsPath == "" {
		c.PgMigrationsPath = defaultPgMigrationsPath
	}

}

func initViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
