package config

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("config/")
	config.AddConfigPath("src/config/")
	fmt.Println(config.ConfigFileUsed())
	err = config.ReadInConfig()
	config.AutomaticEnv()
	if err != nil {
		fmt.Println(err)
		log.Fatal("error on parsing configuration file")
	}

	config.SetConfigType("properties")
	config.SetConfigName(".env")
	config.AddConfigPath("src/")
	config.AddConfigPath(".")
	err = config.MergeInConfig()
	fmt.Println(config.ConfigFileUsed())
	if err != nil {
		//THIS IS NOT FATAL, env file is optional to pass around
		fmt.Println(err)
	}

}

func relativePath(basedir string, path *string) {
	p := *path

	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}
