package main

import (
	"github.com/Deymos01/go-ogen-crud-example/internal/api"
	"github.com/Deymos01/go-ogen-crud-example/internal/oas"
	"log"
	"net/http"
)

func main() {
	handler := api.NewCarHandler()

	router, err := oas.NewServer(handler)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
