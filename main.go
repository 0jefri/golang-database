package main

import (
	"log"

	"github.com/golang-database/database"
)

func main() {
	db, err := database.ConnetDatabase()
	if err != nil {
		log.Fatal("Gagal connet database!!!!", err)
	} else {
		log.Println("Success connet to database !!!")
	}
	defer db.Close()
}
