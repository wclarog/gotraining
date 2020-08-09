package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP_PORT string
	APP       string
	ENV       string
	CSV       Repository
}

type Repository struct {
	CSV_FILENAME string
}

var (
	Values Config
	CSV    Repository
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
	viper.SetDefault("CSV_FILENAME", "")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&CSV)
	fmt.Println("csv")
	fmt.Println(CSV)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	err = viper.Unmarshal(&Values)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	Values.CSV = CSV
}
