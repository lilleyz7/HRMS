package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lillez7/HRMS/internals/database"
)

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fail to load environment variables")
	}
}

func Initialize() {
	LoadDotEnv()
	database.ConnectDB()
}
