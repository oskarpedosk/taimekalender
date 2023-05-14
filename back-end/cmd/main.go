package main

import (
	"log"
	"net/http"
	"taimekalender/back-end/driver"
)

func main() {
	err := driver.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server started at: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", routes()))
}
