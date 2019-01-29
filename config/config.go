package config

import (
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/Unknwon/i18n"
	"github.com/gin-gonic/gin"
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
	App struct {
		Debug  bool   `mapstructure:"debug" validate:"required"`
		Secret string `mapstructure:"secret" validate:"required"`
	} `mapstructure:"app"`
	DB struct {
		Host               string `mapstructure:"host" validate:"required"`
		User               string `mapstructure:"user" validate:"required"`
		Password           string `mapstructure:"password" validate:"required"`
		Name               string `mapstructure:"name" validate:"required"`
		MaxIdleConnections int    `mapstructure:"max_idle_connections" validate:"required,max=20"`
		MaxOpenConnections int    `mapstructure:"max_open_connections" validate:"required,max=50"`
	} `mapstructure:"db"`
	Redis struct {
		Address  string `mapstructure:"address" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		PoolSize int    `mapstructure:"pool_size" validate:"required"`
	} `mapstructure:"redis"`
}

func init() {
	// 去掉烦人的gin提示，在http模块中会根据需要打开
	gin.SetMode(gin.ReleaseMode)

	// load .env file for testing
	godotenv.Load()

	// app
	viper.SetDefault("app.debug", true)
	viper.SetDefault("app.secret", "123456")

	// db
	viper.SetDefault("db.host", "localhost:3306")
	viper.SetDefault("db.user", "user")
	viper.SetDefault("db.password", "password")
	viper.SetDefault("db.name", "name")
	viper.SetDefault("db.max_idle_connections", 15)
	viper.SetDefault("db.max_open_connections", 50)

	// redis
	viper.SetDefault("redis.address", "localhost")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.pool_size", 10)

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
	if C.App.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// load language file for i18n
	files, err := ioutil.ReadDir("languages")
	if err == nil {
		for _, file := range files {
			if err := i18n.SetMessage(strings.TrimSuffix(file.Name(), path.Ext(file.Name())), "languages/"+file.Name()); err != nil {
				log.Fatalf("i18n.SetMessage error: %v\n", err)
			}
		}

		i18n.SetDefaultLang("en-US")
	}
}
