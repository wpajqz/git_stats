package main

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	validator "gopkg.in/go-playground/validator.v9"
)

var c config

type config struct {
	Debug      bool   `mapstructure:"debug" validate:"required"`
	GiteaToken string `mapstructure:"gitea_token" validate:"required"`
	StatsDir   string `mapstructure:"dir" validate:"required"`
	GitDir     string `mapstructure:"git_dir" validate:"required"`
}

func init() {
	// load .env file for testing
	godotenv.Load()

	// app
	viper.SetDefault("debug", true)
	viper.SetDefault("gitea_token", "token")
	viper.SetDefault("dir", "/stats")
	viper.SetDefault("git_dir", "/git")

	// bind env
	viper.SetEnvPrefix("stats")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// unmarshal config to struct
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("viper.Unmarshal error: %v\n", err)
	}

	// validate data type
	validate := validator.New()
	err := validate.Struct(&c)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	// set logrus debug mode
	if c.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
