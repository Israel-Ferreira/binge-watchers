package config

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

var DbConfigVars *DbConfig = &DbConfig{}

func LoadEnv() {
	if err := env.Load(); err != nil {
		log.Println(".env file not found")
	}

	DbConfigVars.Host = os.Getenv("PG_HOST")
	DbConfigVars.Port = os.Getenv("PG_PORT")
	DbConfigVars.Username = os.Getenv("PG_USER")
	DbConfigVars.Password = os.Getenv("PG_PASSWORD")
	DbConfigVars.DbName = os.Getenv("PG_DBNAME")
}
