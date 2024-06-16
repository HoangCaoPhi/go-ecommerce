package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./../../../configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	error := viper.ReadInConfig()
	if error != nil {
		panic(fmt.Errorf("Failed to read configuartion %w \n", error))
	}

	fmt.Println("Server post::", viper.GetInt("server.port"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %e", err)
	}

	fmt.Println("Server post 1::", config.Server.Port)

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	fmt.Println(dir)
}
