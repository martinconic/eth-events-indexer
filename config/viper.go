package config

import "github.com/spf13/viper"

const (
	ConfigName = "config"
	ConfigPath = "."
)

func NewViper() (*viper.Viper, error) {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigName(ConfigName)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}
