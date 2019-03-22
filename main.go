package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

//Configuration is struct for decode..
type Configuration struct {
	App     AppConfiguration
	MongoDb DatabaseConfiguration
	Log     LogConfiguration
}

//AppConfiguration is sub-struct Configuration
type AppConfiguration struct {
	Env   string
	port  int
	debug bool
}

//DatabaseConfiguration is sub-struct Configuration
type DatabaseConfiguration struct {
	Connection string
}

//LogConfiguration is sub-struct Configuration
type LogConfiguration struct {
	Level  string
	Format string
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config Configuration
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Printf("can't decode into struct: %v", err)
	}

	fmt.Printf("%+v \n", config)
	fmt.Printf("%+v \n", config.App)
	fmt.Printf("%+v \n", config.MongoDb)
	fmt.Printf("%+v \n", config.Log)

}
