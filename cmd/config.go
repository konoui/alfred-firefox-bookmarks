package cmd

import (
	"github.com/spf13/viper"
)

// Config configuration which browser bookmark read
type Config struct {
	Firefox         Firefox `json:"firefox"`
	Chrome          Chrome  `json:"chrome"`
	RemoveDuplicate bool    `json:"removeDuplicate"`
}

// Firefox Configuration
type Firefox struct {
	Enable bool   `json:"enable"`
	Path   string `json:"path,omitempty"`
}

// Chrome Configuration
type Chrome struct {
	Enable bool   `json:"enable"`
	Path   string `json:"path,omitempty"`
}

// NewConfig return alfred bookmark configuration
func newConfig() (*Config, error) {
	var c Config
	viper.SetConfigType("yaml")
	viper.SetConfigName(".alfred-bookmarks")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/")

	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return &Config{}, err
	}

	return &c, nil
}
