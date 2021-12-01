package configs

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

type Config struct {
	DbUser 					string
	DbPassword 				string
	DbName 					string
	DriverName 				string
	DbHost    				string
	DbPort    				string
}
const projectDirName = "Dp-218_Go"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/configs/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func Get() *Config {
	loadEnv()

	config := Config {
			DbUser:     os.Getenv("POSTGRES_USER"),
			DbPassword: os.Getenv("POSTGRES_PASSWORD"),
			DbName:     os.Getenv("POSTGRES_DB"),
			DriverName: os.Getenv("DRIVER_NAME"),
			DbHost:     os.Getenv("POSTGRES_HOST"),
			DbPort: 	os.Getenv("POSTGRES_PORT"),
	}
	configBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Configuration:", string(configBytes))

	return &config
}
