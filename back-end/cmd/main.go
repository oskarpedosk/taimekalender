package main

import (
	"github.com/oskarpedosk/taimekalender/back-end/driver"
	"log"
	"net/http"
)

func main() {
	driver.ConnectDB()

	log.Println("Server started at: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", routes()))
}
