package configs

import (
	"github.com/mcuadros/go-defaults"
	"log"
)

import (
	"github.com/spf13/viper"
)

type Configs struct {
	Server Server
	DB     DB
}

func NewServerConfig() *Configs {
	log.Println("Configs is initializing...")
	config := Configs{}
	setDefault(&config)
	viper.AddConfigPath("../../")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viperConfigRead(&config); err != nil {
		log.Fatalf("fatal on change read")
	}

	return &config
}

func viperConfigRead(config *Configs) error {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return nil
}

func setDefault(config *Configs) {
	defaults.SetDefaults(config)
}
