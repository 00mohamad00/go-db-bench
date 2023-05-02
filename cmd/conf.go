package main

import (
	"strings"
	"time"

	"github.com/00mohamad00/go-db-bench/internal/mongo"
	"github.com/00mohamad00/go-db-bench/internal/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Mongo    mongo.Config
	Postgres postgres.Config
}

func LoadConfig(cmd *cobra.Command) (*Config, error) {
	viper.SetDefault("Mongo.URL", "mongodb://localhost:27018")
	viper.SetDefault("Mongo.Database", "test")
	viper.SetDefault("Mongo.Timeout", time.Second)

	viper.SetDefault("Postgres.Host", "localhost")
	viper.SetDefault("Postgres.Port", 5432)
	viper.SetDefault("Postgres.Username", "test")
	viper.SetDefault("Postgres.Password", "test")
	viper.SetDefault("Postgres.DBName", "test")

	// Read Config from ENV
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func loadConfigOrPanic(cmd *cobra.Command) *Config {
	conf, err := LoadConfig(cmd)
	if err != nil {
		logrus.WithError(err).Panic("Failed to load configurations")
	}
	return conf
}
