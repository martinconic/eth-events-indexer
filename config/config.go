package config

import "github.com/spf13/viper"

type PostgresConfig struct {
	host     string
	port     int
	name     string
	user     string
	password string
}

type NetworkConfig struct {
	Key    string
	Wss    string
	Https  string
	Adress string
}

func GetNetworkConfigs(v *viper.Viper) *NetworkConfig {
	return &NetworkConfig{
		Key:   v.GetString("network.key"),
		Wss:   v.GetString("network.wss"),
		Https: v.GetString("network.https"),
	}
}
