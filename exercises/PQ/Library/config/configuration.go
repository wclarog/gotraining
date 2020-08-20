package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP_PORT string
	APP       string
	ENV       string
	DB        Database
}

type Database struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
}

var (
	Values Config
	DB     Database
)

func init() {
	// Config file extension
	viper.SetConfigType("json")
	// Include env vars (Env vars precedence is bigger than config files)
	viper.AutomaticEnv()
	// Config file name
	viper.SetConfigName("config")
	// Config path
	viper.AddConfigPath("./config")
	// If values are loaded from env vars only set defaults
	viper.SetDefault("DB_USER", "")
	viper.SetDefault("DB_PASS", "")
	viper.SetDefault("DB_HOST", "")
	viper.SetDefault("DB_PORT", "")
	viper.SetDefault("DB_NAME", "")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&DB)
	fmt.Println("db")
	fmt.Println(DB)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	err = viper.Unmarshal(&Values)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	Values.DB = DB
}
