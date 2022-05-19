package main

import (
	"errors"
	"fngc/models"
	"fngc/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// @title Swift medics Backend API Service
// @version 1.0
// @description This is official backend documentation for the swift medics backend system

// @contact.name Endy Apinageri
// @contact.email apinaendy@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1
func main() {
	err := godotenv.Load("conf.env")
	if err != nil {
		panic("Error loading .env file")
	}

	appPort := os.Getenv("port")
	hostAddress := os.Getenv("host_address")
	if os.Getenv("app_mode") == "prod" {
		hostAddress = os.Getenv("prod_host_address")
	}
	log.Println("App running on " + hostAddress + appPort)

	r := routes.SetupRouter(appPort, hostAddress)

	if os.Getenv("app_mode") == "dev" {
		http.ListenAndServe(appPort, r)
	} //run app on development mode

	if os.Getenv("app_mode") == "prod" {
		HTTPSCertFile := "/etc/letsencrypt/live/www.gasnigeriaapi.com/fullchain.pem"
		HTTPSKeyFile := "/etc/letsencrypt/live/www.gasnigeriaapi.com/privkey.pem"
		if err := http.ListenAndServeTLS(appPort, HTTPSCertFile, HTTPSKeyFile, r); err != nil {
			models.LogError(errors.New("error starting test-conf server"))
			models.LogError(err)
		}
	} //run app on production mode
}
