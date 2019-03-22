package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println(viper.Get("app.env"))
	fmt.Println(viper.Get("app.port"))
	fmt.Println(viper.Get("app.debug"))
	fmt.Println("=========================")

	fmt.Println(viper.GetString("mongodb.connection"))
	fmt.Printf("%T \n", viper.GetString("mongodb.connection"))
	fmt.Println("=========================")

	fmt.Println(viper.GetBool("app.debug"))
	fmt.Println("=========================")

	if viper.IsSet("app.debug") {
		fmt.Println("Yep! debug")
	}
	fmt.Println("=========================")

	log := viper.GetStringMap("log")
	fmt.Println(log["level"])
	fmt.Println(log["format"])
	fmt.Println("=========================")

	viper.SetDefault("app.env", "production")
	fmt.Println(viper.Get("app.env"))

}
