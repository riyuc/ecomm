package main

import (
	"log"

	"github.com/riyuc/ecomm/db"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error opening database: %v",err)
	}
	defer db.Close()
	log.Println("Connected to database")
}