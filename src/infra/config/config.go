package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	UrlValidator = ""
	ApiPort      = 0
	Host         = ""
	XApiKey      = ""
	XApiGtwId    = ""
)

func VariablesLoading() {
	var err error

	if godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ApiPort = 8082
	}

	Host = os.Getenv("HOST")
	XApiKey = os.Getenv("XAPIKEY")
	XApiGtwId = os.Getenv("XAPIGTWID")
}
