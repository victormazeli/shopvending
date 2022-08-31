package cmd

import "github.com/spf13/viper"

type ServerConfig struct {
	Port       string `mapstructure:"env:PORT"`
	Name       string `mapstructure:"env:SERVICE_NAME"`
	Host       string `mapstructure:"env:DATABASE_Host"`
	DBPassword string `mapstructure:"env:DATABASE_PASSWORD"`
	User       string `mapstructure:"env:DATABASE_USER"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config ServerConfig, err error) {
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
