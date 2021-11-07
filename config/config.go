package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

// Config of app
type Config struct {
	JSON *viper.Viper

	HTTPAddr            string `envconfig:"HTTP_ADDR"`
	URIScheme           string `envconfig:"URI_SCHEME"`
	CorsAllowedAddr     string `envconfig:"CORS_ALLOWED_ADDR"`
	LogLevel            string `envconfig:"LOG_LEVEL"`
	PgURL               string `envconfig:"PG_URL"`
	REDIS_URL           string `envconfig:"REDIS_URL"`
	PgMigrationsPath    string `envconfig:"PG_MIGRATIONS_PATH"`
	MysqlAddr           string `envconfig:"MYSQL_ADDR"`
	MysqlUser           string `envconfig:"MYSQL_USER"`
	MysqlPassword       string `envconfig:"MYSQL_PASSWORD"`
	MysqlDB             string `envconfig:"MYSQL_DB"`
	MysqlMigrationsPath string `envconfig:"MYSQL_MIGRATIONS_PATH"`
	GCBucket            string `envconfig:"GC_BUCKET"`
	JobsPeriod          int    `envconfig:"JOBS_PERIOD"`
	SessionKey          string `envconfig:"SESSION_KEY"`
	TransactionSsupport bool   `envconfig:"TRANSACTION_SUPPORT"`
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
		fmt.Println("Configuration:", string(configBytes))
		config.JSON = viper.GetViper()
	})
	return &config
}

func initViper() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}