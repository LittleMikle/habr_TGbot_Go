package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramBotToken string `mapstructure:"telegramBotToken"`
	WebhookURL       string `mapstructure:"WebhookURL"`
	Port             string `mapstructure:"port"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) Load(path string, name string, _type string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(_type)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read config error: %w", err)
	}

	err = viper.Unmarshal(c)

	if err != nil {
		return fmt.Errorf("unmarshalling config error: %w", err)
	}
	return nil
}
