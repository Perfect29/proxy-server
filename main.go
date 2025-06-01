package main

import (
	"net/http"
	"log"
	"github.com/Perfect29/proxy-server/handlers"
)

func main() {
	http.HandleFunc("/proxy", handlers.HandleProxyRequest)
	http.HandleFunc("/logs/", handlers.HandleGetLog)
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
