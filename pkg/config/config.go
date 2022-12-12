package config

import "github.com/spf13/viper"

type Config struct {
	PostgresURL  string `mapstructure:"POSTGRES_URL"`
	StartParse int `mapstructure:"START_PARSE"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}