package main

import (
	"context"
	"fmt"
	"github.com/Deymos01/go-ogen-crud-example/internal/oas"
	"log"
	"net/http"
)

func main() {
	client, err := oas.NewClient("http://localhost:8080", oas.WithClient(http.DefaultClient))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Add car
	newCar := oas.NewCar{
		Manufacturer: "Tesla",
		Model:        "Model S",
		Year:         2024,
		Color:        "White",
	}
	createdCar, err := client.AddCar(ctx, &newCar)
	if err != nil {
		log.Fatalf("failed to add car: %v", err)
	}
	fmt.Printf("Created car: %+v\n", createdCar)

	// Get car by ID
	car, err := client.GetCarById(ctx, oas.GetCarByIdParams{ID: 0})
	if err != nil {
		log.Fatalf("failed to get car: %v", err)
	}
	fmt.Printf("Fetched car: %+v\n", car)

	// Update car by ID
	updCar := oas.UpdateCar{
		Model: oas.OptString{Set: true, Value: "Model X"},
		Color: oas.OptString{Set: true, Value: "Black"},
	}
	updRes, err := client.UpdateCarById(ctx, &updCar, oas.UpdateCarByIdParams{ID: 0})
	if err != nil {
		log.Fatalf("failed to update car: %v", err)
	}
	fmt.Printf("Updated car: %+v\n", updRes)

	// Delete car by ID
	delRes, err := client.DeleteCarById(ctx, oas.DeleteCarByIdParams{ID: 0})
	if err != nil {
		log.Fatalf("failed to delete car: %v", err)
	}
	fmt.Printf("Delete result: %T\n", delRes)
}
