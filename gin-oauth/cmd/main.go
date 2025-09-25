package main

import (
	"fmt"
	"go-oauth2/internal/adapter/config"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Configuration imported successfully ✅")

	fmt.Println(conf.App)
}
