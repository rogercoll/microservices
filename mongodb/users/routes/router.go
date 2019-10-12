package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes() router *mux.Router {
	router := mux.NewRouter()
	router = UserRoutes(router)
}
