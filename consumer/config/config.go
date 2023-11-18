package config

import "github.com/spf13/viper"

type EnvConfig struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	AppName       string `mapstructure:"APP_NAME"`
	AppPort       string `mapstructure:"APP_PORT"`
	AppVersion    string `mapstructure:"APP_VERSION"`
	KafkaSasl     bool   `mapstructure:"KAFKA_SASL"`
	KafkaHosts    string `mapstructure:"KAFKA_HOSTS"`
	KafkaUsername string `mapstructure:"KAFKA_USERNAME"`
	KafkaPassword string `mapstructure:"KAFKA_PASSWORD"`
}

func Load(filename string) (*EnvConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var envCfg EnvConfig
	if err := viper.Unmarshal(&envCfg); err != nil {
		return nil, err
	}

	return &envCfg, nil
}
