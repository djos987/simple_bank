package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER" json:"db_driver" yaml:"db_driver"`
	DBSource      string `mapstructure:"DB_SOURCE" json:"db_source" yaml:"db_source"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS" json:"server_address" yaml:"server_address"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
