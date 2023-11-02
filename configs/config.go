package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

func InConfig() {
	// Load configuration from config.yaml
	viper.SetConfigFile("configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}
}
