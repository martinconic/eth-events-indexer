package config

import (
	"github.com/martinconic/eth-events-indexer/utils/constants"
	"github.com/spf13/viper"
)

type PostgresConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type NetworkConfig struct {
	Key    string
	Wss    string
	Https  string
	Adress string
}

func GetNetworkConfigs(v *viper.Viper) *NetworkConfig {
	return &NetworkConfig{
		Key:   v.GetString(constants.NetworkKey),
		Wss:   v.GetString(constants.NetworkWss),
		Https: v.GetString(constants.NetworkHttps),
	}
}

func GetPostgresConfig(v *viper.Viper) *PostgresConfig {
	return &PostgresConfig{
		Host:     v.GetString(constants.PostgresHost),
		Port:     v.GetInt(constants.PostgresPort),
		Name:     v.GetString(constants.PostgresName),
		User:     v.GetString(constants.PostgresUser),
		Password: v.GetString(constants.PostgresPass),
	}
}
