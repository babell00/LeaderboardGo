package server

import (
	"log"
	"fmt"
	"net/http"
)

func NewServer(portNumber int) {
	log.Println("Setting up Server")

	router := NewRouter()
	formattedPortNumber := formatPortNumber(portNumber)

	log.Printf("Server is listening on port: %v", portNumber)
	log.Fatal("Server error: ", http.ListenAndServe(formattedPortNumber, router))

}


func formatPortNumber(portNumber int) string {
	return fmt.Sprintf(":%v", portNumber)
}