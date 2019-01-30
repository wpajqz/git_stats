package config

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	validator "gopkg.in/go-playground/validator.v9"
)

// global config
var (
	C appConfig
)

type appConfig struct {
	Debug      bool   `mapstructure:"debug" validate:"required"`
	GiteaToken string `mapstructure:"gitea_token" validate:"required"`
	StatsDir   string `mapstructure:"dir" validate:"required"`
}

func init() {
	// load .env file for testing
	godotenv.Load()

	// app
	viper.SetDefault("debug", true)
	viper.SetDefault("gitea_token", "token")
	viper.SetDefault("dir", "/stats")

	// bind env
	viper.SetEnvPrefix("stats")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// unmarshal config to struct
	if err := viper.Unmarshal(&C); err != nil {
		log.Fatalf("viper.Unmarshal error: %v\n", err)
	}

	// validate data type
	validate := validator.New()
	err := validate.Struct(&C)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	// set logrus debug mode
	if C.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
