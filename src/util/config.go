package util

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string         `mapstructure:"environment"`
	Database    DatabaseConfig `mapstructure:"database"`
	Server      ServerConfig   `mapstructure:"server"`
	Token       TokenConfig    `mapstructure:"token"`
	Email       EmailConfig    `mapstructure:"email"`
}

type DatabaseConfig struct {
	Driver     string `mapstructure:"driver"`
	Engine     string `mapstructure:"engine"`
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Name       string `mapstructure:"name"`
	SSLMode    string `mapstructure:"sslmode"`
	MigrateUrl string `mapstructure:"migration_url"`
}

type ServerConfig struct {
	HTTPAddress  string `mapstructure:"http_address"`
	GRPCAddress  string `mapstructure:"grpc_address"`
	RedisAddress string `mapstructure:"redis_address"`
}

type TokenConfig struct {
	SymetricKey          string        `mapstructure:"symetric_key"`
	AccessTokenDuration  time.Duration `mapstructure:"access_token_duration"`
	RefreshTokenDuration time.Duration `mapstructure:"refresh_token_duration"`
}

type EmailConfig struct {
	SenderName     string `mapstructure:"sender_name"`
	SenderAddress  string `mapstructure:"sender_address"`
	SenderPassword string `mapstructure:"sender_password"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func (c *DatabaseConfig) GetDBSource() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", c.Engine, c.Username, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
}
