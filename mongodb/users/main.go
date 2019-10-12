package main

import (
	"log"
	"net/http"
	"github.com/rcoll/microservices/mongodb/users/config"
	"github.com/rcoll/microservices/mongodb/users/routers"
)

func main(){
	config.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	server := &http.Server{
		Addr:    config.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}