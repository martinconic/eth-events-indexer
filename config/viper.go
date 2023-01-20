package config

import "github.com/spf13/viper"

const (
	ConfigName = ".env"
	ConfigPath = "."
)

func NewViper() (*viper.Viper, error) {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigFile(ConfigName)
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}
