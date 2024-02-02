package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App    App
		Server Server
		Db     Db
	}

	App struct {
		Profile string
		XxxUrl string
	}

	Server struct {
		Addr string
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}
)

func Load() *Config {
	viper.AutomaticEnv() //auto read key env if yaml value is empty
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}

	return &Config{
		App: App{
			Profile: viper.GetString("app.profile"),
			XxxUrl:  viper.GetString("app.xxxurl"),
		},
		Server: Server{
			Addr: viper.GetString("server.addr"),
		},
		Db: Db{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			DBName:   viper.GetString("database.dbname"),
			SSLMode:  viper.GetString("database.sslmode"),
			TimeZone: viper.GetString("database.timezone"),
		},
	}
}
