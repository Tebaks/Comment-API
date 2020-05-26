package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User     string
		Password string
		Net      string
		Addr     string
		DBName   string
	}
	Server struct {
		Address string
	}
}

var C config

func ReadConfig() {
	Config := &C
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	wd, _ := os.Getwd()
	viper.AddConfigPath(filepath.Join(wd, "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
