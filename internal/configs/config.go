package configs

import "github.com/spf13/viper"

type Config struct {
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_HOST"`
	PostgresUser     string `mapstructure:"POSTGRES_HOST"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresSslMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	ClusterId        string `mapstructure:"CLUSTER_ID"`
	ClientProducer   string `mapstructure:"CLIENT_PRODUCER"`
	ClientSubscriber string `mapstructure:"CLIENT_SUBSCRIBER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
