package main

import (
	"authentication/routes"
	"authentication/internal/config"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

func main() {
	log.Println("Starting Authentication Service")

	// TODO: Connect to db

	// set up config
	app := config.Config{}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: routes.Routes(app),
	}

	err:= srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}