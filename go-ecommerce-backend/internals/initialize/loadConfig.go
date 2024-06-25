package initialize

import (
	"fmt"

	"github.com/HoangCaoPhi/go-ecommerce/globals"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()

	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	if err := viper.Unmarshal(&globals.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
