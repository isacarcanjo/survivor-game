package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               int
	ConnectionString   string
	Env                string
	Version            string
	SurvivorCollection string
}

func GetDotEnv() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	Port, err2 := strconv.Atoi(os.Getenv("PORT"))
	if err2 != nil {
		Port = 5000
	}
	ConfigData := Config{}
	ConfigData.Port = Port
	ConfigData.Version = os.Getenv("VERSION")
	ConfigData.ConnectionString = fmt.Sprintf("%s:%s@/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	ConfigData.Env = os.Getenv("ENV")
	ConfigData.SurvivorCollection = os.Getenv("SURVIVOR_COLLECTION")
	return ConfigData
}
