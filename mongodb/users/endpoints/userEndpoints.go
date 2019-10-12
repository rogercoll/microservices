package endpoints

import (
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/mux"
	"github.com/rogercoll/microservices/mongodb/users/data"
	"github.com/rogercoll/microservices/mongodb/users/config"

)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	users := repo.GetAll()
	j, err := json.Marshal(UsersResource{Data: users})
	if err != nil {
		config.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		config.DisplayAppError(w, err, "Invalid User data", 500)
		return
	}
	user := &dataResource.Data
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	// Create User
	repo := &data.UserRepository{c}
	repo.Create(user)
	// Create response data
	j, err := json.Marshal(dataResource)
	if err != nil {
		config.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")

	// Remove user by id
	repo := &data.UserRepository{c}
	err := repo.Delete(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			config.DisplayAppError(w, err, "An unexpected error ahs occurred", 500)
			return
		}
	}

	// Send response back
	w.WriteHeader(http.StatusNoContent)
}