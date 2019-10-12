package main

import (
	"log"
	"net/http"
	"github.com/rogercoll/microservices/mongodb/users/config"
	"github.com/rogercoll/microservices/mongodb/users/routes"
)


func main() {
	config.StartUp()
	router := routes.InitRoutes()
	server := &http.Server{
		Addr:    config.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}