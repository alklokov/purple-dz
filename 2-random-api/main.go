package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", randomHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

func randomHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(rand.Intn(6) + 1)))
}
