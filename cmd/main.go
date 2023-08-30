package main

import (
	"bmc/bmc-assignment/cmd/api"
	"log"

	_ "bmc/bmc-assignment/docs"
)

// @title BMC-assignment 2023-08-29
// @version 1.0
// @description BMC Histogram CSV SQLite Docker

// @host localhost:8080
// @BasePath /
func main() {
	err := api.StartApp()
	if err != nil {
		log.Fatalln(err)
	}
}
