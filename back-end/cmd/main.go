package main

import (
	"log"
	"net/http"
	"taimekalender/back-end/driver"
)

func main() {
	db, err := driver.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Server started at: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", routes()))
}
