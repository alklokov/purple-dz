package main

import (
	"3-validation-api/configs"
	"3-validation-api/internal/mailer"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)

	router := http.NewServeMux()
	mailer.MakeMailHandler(router, conf)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
