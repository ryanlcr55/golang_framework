package configs

import (
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
	"go_framework/internal/configs"
	"log"
)

func NewServerConfig() *configs.Configs {
	log.Println("Configs is initializing...")
	config := configs.Configs{}
	setDefault(&config)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viperConfigRead(&config); err != nil {
		log.Fatalf("fatal on change read")
	}

	return &config
}

func viperConfigRead(config *configs.Configs) error {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return nil
}

func setDefault(config *configs.Configs) {
	defaults.SetDefaults(config)
}
