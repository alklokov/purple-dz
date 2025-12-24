package main

import (
	"4-order-api/configs"
	"4-order-api/pkg/db"
	"fmt"
)

func main() {
	conf := configs.LoadConfig()
	myDb := db.OpenDb(conf)
	fmt.Println("DB is Open: ", myDb)
}
