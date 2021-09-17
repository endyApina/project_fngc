package main

import (
	"fngc/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// @title FNGC Backend Service
// @version 1.0
// @description This is official backend documentation for the fngc software

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
	log.Println("App running on " + hostAddress + appPort)

	r := routes.SetupRouter(appPort, hostAddress)
	http.ListenAndServe(appPort, r)
}
