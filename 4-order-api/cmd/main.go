package main

import (
	"4-order-api/configs"
	"4-order-api/internal/product"
	"4-order-api/pkg/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	myDb := db.OpenDb(conf)
	fmt.Println("DB is Open: ", myDb)

	router := http.NewServeMux()
	port := 8081

	productRepository := product.NewProductRepository(myDb)
	product.RegisterProductHandler(router, productRepository)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	fmt.Printf("Server started on port %d\n", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error", err)
	}
}
