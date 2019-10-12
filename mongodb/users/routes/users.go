package routes

import (
	"github.com/gorilla/mux"
	"github.com/rcoll/microservices/mongodb/users/endpoints"
)

func UserRoutes(router *mux.Router) router *mux.Router {
	router.HandleFunc("/users", endpoints.GetUsers).Methods("GET")
	router.HandleFunc("/users", endpoints.InsertUser).Methods("POST")
	router.HandleFunc("/users/{id}", endpoints.DeleteUser).Methods("DELETE")
}