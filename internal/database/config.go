package database

import "github.com/spf13/viper"

type Config struct {
	Host             string
	Port             int
	Username         string
	Password         string
	Database         string
	SSL              string
	MigrationVersion int
}

func LoadConfig() Config {
	viper.AutomaticEnv()
	return Config{
		Host:             viper.GetString("POSTGRES_HOST"),
		Port:             viper.GetInt("POSTGRES_PORT"),
		Database:         viper.GetString("POSTGRES_DB"),
		SSL:              viper.GetString("POSTGRES_SSL"),
		Username:         viper.GetString("POSTGRES_USER"),
		Password:         viper.GetString("POSTGRES_PASSWORD"),
		MigrationVersion: viper.GetInt("POSTGRES_MIGRATIONVERSION"),
	}
}
