// Package config provides configuration loading and validation.
package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config holds all application configuration.
type Config struct {
	App         AppConfig         `mapstructure:"app"`
	TigerBeetle TigerBeetleConfig `mapstructure:"tigerbeetle"`
}

// AppConfig holds general application settings.
type AppConfig struct {
	Name     string `mapstructure:"name"`
	LogLevel string `mapstructure:"log_level"`
}

// TigerBeetleConfig holds TigerBeetle connection settings.
type TigerBeetleConfig struct {
	ClusterID      uint128String `mapstructure:"cluster_id"`
	Addresses      []string      `mapstructure:"addresses"`
	MaxConcurrency uint          `mapstructure:"max_concurrency"`
	ConnectTimeout time.Duration `mapstructure:"connect_timeout"`
}

// uint128String is a string representation of a uint128 cluster ID.
type uint128String = string

// Load loads configuration from file and environment variables.
func Load(configPath string) (*Config, error) {
	v := viper.New()

	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("./config")
	}

	v.SetEnvPrefix("TIGER")
	v.AutomaticEnv()

	bindEnvVars(v)
	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &cfg, nil
}

func bindEnvVars(v *viper.Viper) {
	v.BindEnv("app.name", "TIGER_APP_NAME")
	v.BindEnv("app.log_level", "TIGER_LOG_LEVEL", "LOG_LEVEL")
	v.BindEnv("tigerbeetle.cluster_id", "TIGER_TB_CLUSTER_ID", "TB_CLUSTER_ID")
	v.BindEnv("tigerbeetle.addresses", "TIGER_TB_ADDRESSES", "TB_ADDRESSES")
	v.BindEnv("tigerbeetle.max_concurrency", "TIGER_TB_MAX_CONCURRENCY")
	v.BindEnv("tigerbeetle.connect_timeout", "TIGER_TB_CONNECT_TIMEOUT")
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.name", "tiger-tui")
	v.SetDefault("app.log_level", "info")
	v.SetDefault("tigerbeetle.cluster_id", "0")
	v.SetDefault("tigerbeetle.addresses", []string{"3000"})
	v.SetDefault("tigerbeetle.max_concurrency", 32)
	v.SetDefault("tigerbeetle.connect_timeout", "5s")
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	if len(c.TigerBeetle.Addresses) == 0 {
		return fmt.Errorf("tigerbeetle.addresses cannot be empty")
	}
	return nil
}
