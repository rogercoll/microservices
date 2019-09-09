package main

import (
	"time"
	"context"
    "flag"
	"fmt"
	"io/ioutil"
    "log"
    "net/http"
	"os"
	"bytes"
    "os/signal"
	"syscall"
    "encoding/json"
	"github.com/rogercoll/microservices/datastoregcp"
)

type SomethingCool struct {
	ProjID	string	`json:"projID"`
	EntityName	string `json:"entityName"`
}


func testRequest() {
	url := "http://localhost:8081/getObject"	
	test := SomethingCool{EntityName: "Book", ProjID: "rcoll-laboratorio"}
	b, err := json.Marshal(test)
	fmt.Println(string(b))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

func main() {
	var (
		httpAddr = flag.String("http", ":8081", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := datastoregcp.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoints := datastoregcp.Endpoints {
		GetEndpoint: 	datastoregcp.MakeGetEndpoint(srv),
		StoreEndpoint:	datastoregcp.MakeStoreEndpoint(srv),
	}

	go func() {
		log.Println("DataStoreGcp microservice is listening on port: ",*httpAddr)
		handler := datastoregcp.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	time.Sleep(3*time.Second)
	testRequest()
	log.Fatal(<-errChan)
}