package config

import (
	"github.com/mohammadVatandoost/ingbusiness/internal/services/authentication"
	"strings"

	"github.com/mohammadVatandoost/ingbusiness/internal/core/grpc"

	"github.com/mohammadVatandoost/ingbusiness/internal/database"
	"github.com/mohammadVatandoost/ingbusiness/internal/goadmin"
	"github.com/mohammadVatandoost/ingbusiness/pkg/logger"
	"github.com/mohammadVatandoost/ingbusiness/pkg/prometheus"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config the application's configuration structure
type Config struct {
	ConfigFile string
	Logger     logger.Config
	GRPC       grpc.Config
	Postgres   database.Config
	Metric     prometheus.Config
	GoAdmin    goadmin.Config
	Auth       authentication.Config
}

// LoadConfig loads the config from a file if specified, otherwise from the environment
func LoadConfig(cmd *cobra.Command) (*Config, error) {
	// Setting defaults for this application

	viper.SetDefault("metric.ListenPort", 9000)

	viper.SetDefault("logger.SentryEnabled", false)
	viper.SetDefault("logger.level", "error")

	viper.SetDefault("grpc.ListenPort", 8888)
	viper.SetDefault("grpc.TimeOut", 5)

	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", 5432)
	viper.SetDefault("postgres.username", "postgres")
	viper.SetDefault("postgres.password", "123")
	viper.SetDefault("postgres.database", "postgres")
	viper.SetDefault("postgres.ssl", "disable")
	viper.SetDefault("postgres.MigrationVersion", 1)

	viper.SetDefault("goadmin.port", 9027)
	viper.SetDefault("goadmin.host", "localhost")

	viper.SetDefault("auth.JwtSecret", "testKey")
	viper.SetDefault("auth.GoogleKey", "GoogleKey")
	viper.SetDefault("auth.GoogleSecret", "GoogleSecret")

	// Read Config from ENV
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Read Config from Flags
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Read Config from file
	if configFile, err := cmd.Flags().GetString("config-file"); err == nil && configFile != "" {
		viper.SetConfigFile(configFile)

		if err := viper.ReadInConfig(); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &config, nil
}

func LoadTestConfig() (*Config, error) {

	viper.SetDefault("metric.ListenPort", 9000)

	viper.SetDefault("logger.SentryEnabled", false)
	viper.SetDefault("logger.level", "error")

	viper.SetDefault("grpc.ListenPort", 8888)
	viper.SetDefault("grpc.TimeOut", 5)

	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", 5432)
	viper.SetDefault("postgres.username", "root")
	viper.SetDefault("postgres.password", "root")
	viper.SetDefault("postgres.database", "test_db")
	viper.SetDefault("postgres.ssl", "disable")
	viper.SetDefault("postgres.MigrationVersion", 1)

	viper.SetDefault("goadmin.port", 9027)
	viper.SetDefault("goadmin.host", "localhost")

	// Read Config from ENV
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &config, nil
}
