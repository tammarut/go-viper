package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	//=>ชื่อfile config.yaml
	viper.SetConfigName("config")

	//=>Path
	viper.AddConfigPath("..")

	//=>อ่านvalue จากENV variable
	viper.AutomaticEnv()

	//=> แทน _ ด้วย . ในenv
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	//=> read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//.Viper Method
	fmt.Println(viper.Get("app.env"))
	fmt.Println(viper.Get("app.port"))
	fmt.Println(viper.Get("app.debug"))
	fmt.Println("=========================")

	//=>Get as "string"
	fmt.Println(viper.GetString("mongodb.connection"))
	fmt.Printf("%T \n", viper.GetString("mongodb.connection"))
	fmt.Println("=========================")

	//=>Get as bool(true/false)
	fmt.Println(viper.GetBool("app.debug"))
	fmt.Println("=========================")

	//=>Return bool
	if viper.IsSet("app.debug") {
		fmt.Println("Yep! debug")
	}
	fmt.Println("=========================")

	//=>String Map
	log := viper.GetStringMap("log")
	fmt.Println(log["level"])
	fmt.Println(log["format"])
	fmt.Println(log["formats"])
	fmt.Printf("%T\n", log["level"])
	fmt.Println("=========================")

	//=>as "String" inception
	actionCode := "acm024-001"
	url := viper.GetString("actioncode." + actionCode + ".url")
	fmt.Println(url)

	fmt.Println("=========================")
	//=>String Map inception
	urlmap := viper.GetStringMap("actioncode")
	fmt.Println(urlmap["url"+actionCode])
	fmt.Printf("%T\n", urlmap["url"+actionCode])

	fmt.Println("=========================")
	//=>String Map String
	urlst := viper.GetStringMapString("actioncode")
	fmt.Println(urlst["url"+actionCode])
	fmt.Printf("%T\n", urlst["url"+actionCode])

	fmt.Println("=========================")

	//=>Set default env
	viper.SetDefault("app.env", "production")
	fmt.Println(viper.Get("app.env"))

}
