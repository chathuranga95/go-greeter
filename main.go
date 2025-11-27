package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Response structure for JSON format
type Response struct {
	Hello   string `json:"hello"`
	Port    int    `json:"port"`
	Version string `json:"version"`
}

// Handler for the first service on port 9090
func handler(w http.ResponseWriter, r *http.Request) {
	response := Response{Hello: "world", Port: 9090, Version: "v1.0"}
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Greet service is called at %s\n",
		time.Now().Format(time.RFC3339))
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Setting up the first service listener on port 9090
	http.HandleFunc("/greet", handler)
	go func() {
		if err := http.ListenAndServe(":9090", nil); err != nil {
			log.Printf("Error starting server on port 9090: %v\n", err)
		}
		log.Println("Server started on port 9090")
	}()

	// Prevent the main function from exiting
	select {}
}
