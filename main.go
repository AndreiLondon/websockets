package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	setupAPI()

	// Serve on port :8080, fudge yeah hardcoded port
	fmt.Println("Application started at port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

// setupAPI will start all Routes and their Handlers

func setupAPI() {
	// Create a Manager instance used to handle WebSocket Connections
	manager := NewManager()

	// Serve the ./frontend directory at Route /
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
}
