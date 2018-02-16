package main

import (
	"github.com/spf13/viper"
)

var config *viper.Viper

func initConfig()  {
	config = viper.New()
	config.SetConfigFile("config/app.yml")

	err := config.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Fatal().Msgf("Fatal error config file: %s", err)
	}
}

func saveConfig()  {
	err := config.WriteConfigAs("config/example.yml")
	if err != nil {
		log.Error().Err(err).Msg("saveConfig error")
	}
}