package main

import (
	"log"
	"net/http"
	"taimekalender/back-end/driver"
)

func main() {
	driver.ConnectDB()

	log.Println("Server started at: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", routes()))
}
