package main

import (
	"log"

	"github.com/datrine/conn"
	"github.com/datrine/routes"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	db := conn.InitGetDB()
	if db != nil {
		println("Got connected.")
	}
}

func main() {
	r := routes.SetupRouters()
	// Listen and Server in 0.0.0.0:7000
	r.Run(":7000")
}
