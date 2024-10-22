package config

import (
	"github.com/spf13/viper"
)

const (
	EnvKey        = "ENVIRONMENT"
	ChatGPTApiKey = "CHAT_API_KEY"
	GeminiAPIKey  = "GEMINI_API_KEY"
)

type Config struct {
	ChatGPTApiKey string `mapstructure:"CHAT_API_KEY"`
	GeminiAPIKey  string `mapstructure:"GEMINI_API_KEY"`
}

var config Config

func GetCurrentConfig() Config {
	return config
}

func GetGeminiAPIKey() string {
	return config.GeminiAPIKey
}

func GetEnv() string {
	return viper.GetString(EnvKey)
}

func LoadConfigs(path string) (Config, error) {

	var (
		err error
	)

	viper.AddConfigPath(path)
	viper.SetConfigName("")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if viper.GetString(EnvKey) != "prod" {

		err = viper.ReadInConfig()
		if err != nil {
			return config, err
		}

		err = viper.Unmarshal(&config)
		if err != nil {
			return config, err
		}

		return config, err
	} else {

		config.ChatGPTApiKey = viper.GetString(ChatGPTApiKey)
		config.GeminiAPIKey = viper.GetString(GeminiAPIKey)

		return config, err

	}

}
