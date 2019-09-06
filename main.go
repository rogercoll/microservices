package main

import (
	"context"
	"flag"
	"github.com/rogercoll/microservices/napodate"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := napodate.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}

	endpoints := napodate.Endpoints {
		GetEndpoint: 	napodate.MakeGetEndpoint(srv),
		StatusEndpoint:	napodate.MakeStatusEndpoint(srv),
		ValidateEndpoint:	napodate.MakeValidateEndpoint(srv),
	}

	go func() {
		log.Println("Napodate microservice is listening on port: "*httpAddr)
		handler := napodate.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}

	log.Fataln(<-errChan)
}