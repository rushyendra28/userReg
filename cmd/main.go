package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"honnef.co/go/tools/config"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	} else {
		fmt.Println("Environment variables loaded successfully")
	}

	var conf config.Config
	if err := env.Load(); err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}
}
