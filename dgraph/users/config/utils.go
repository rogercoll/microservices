package config


import (
	"os"
	"log"
	"net/http"
	"encoding/json"
	"google.golang.org/grpc"
	"github.com/dgraph-io/dgo/v2"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	configuration struct {
		Server, DgraphHost, DBUser, DBPwd, Database string
	}
)

var AppConfig configuration
var session *dgo.Dgraph

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error: handlerError.Error(),
		Message: message,
		HttpStatus: code,
	}
	log.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

func initConfig() {
	file, err := os.Open("config/config.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSession() *dgo.Dgraph {
	if session == nil {
		var err error
		conn, err = grpc.Dial(
		AppConfig.DgraphHost,
		grpc.withInsecure()
	)
	if err != nil {
		log.Fatal(err)
	}
	session = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	}
	return session
}

func createDbSession(){
	var err error
	conn, err = grpc.Dial(
		AppConfig.DgraphHost,
		grpc.withInsecure()
	)
	if err != nil {
		log.Fatal(err)
	}
	session = dgo.NewDgraphClient(api.NewDgraphClient(conn))
}