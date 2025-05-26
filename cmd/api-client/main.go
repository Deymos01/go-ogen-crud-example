package main

import (
	"context"
	"flag"
	"github.com/Deymos01/go-ogen-crud-example/internal/oas"
	"github.com/go-faster/errors"
	"log"
	"net/http"
)

func createCar(ctx context.Context, client *oas.Client, newCar *oas.NewCar) error {
	res, err := client.AddCar(ctx, newCar)
	if err != nil {
		return errors.Wrap(err, "create car")
	}
	log.Printf("Created car: %+v\n", res)
	return nil
}

func fetchCar(ctx context.Context, client *oas.Client, id int) error {
	res, err := client.GetCarById(ctx, oas.GetCarByIdParams{ID: id})
	if err != nil {
		return errors.Wrap(err, "get car")
	}
	log.Printf("Got car: %+v\n", res)
	return nil
}

func updateCar(ctx context.Context, client *oas.Client, id int, car *oas.UpdateCar) error {
	res, err := client.UpdateCarById(ctx, car, oas.UpdateCarByIdParams{ID: id})
	if err != nil {
		return errors.Wrap(err, "update car")
	}
	log.Printf("Updated car: %+v\n", res)
	return nil
}

func deleteCar(ctx context.Context, client *oas.Client, id int) error {
	res, err := client.DeleteCarById(ctx, oas.DeleteCarByIdParams{ID: id})
	if err != nil {
		return errors.Wrap(err, "delete car")
	}
	log.Printf("Deleted car: %+v\n", res)
	return nil
}

func main() {
	var err error
	var arg struct {
		BaseURL string
		ID      int
	}
	flag.StringVar(&arg.BaseURL, "url", "http://localhost:8080", "target server url")
	flag.IntVar(&arg.ID, "id", 0, "car id to request")
	flag.Parse()

	client, err := oas.NewClient(arg.BaseURL, oas.WithClient(http.DefaultClient))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Add car
	err = createCar(ctx, client, &oas.NewCar{Manufacturer: "Tesla", Model: "Model S", Year: 2024, Color: "White"})
	if err != nil {
		log.Fatal(err)
	}

	// Get car by ID
	err = fetchCar(ctx, client, arg.ID)
	if err != nil {
		log.Fatal(err)
	}

	// Update car by ID
	err = updateCar(ctx, client, arg.ID, &oas.UpdateCar{
		Model: oas.OptString{Set: true, Value: "Model X"},
		Color: oas.OptString{Set: true, Value: "Black"}})
	if err != nil {
		log.Fatal(err)
	}

	// Delete car by ID
	err = deleteCar(ctx, client, arg.ID)
	if err != nil {
		log.Fatal(err)
	}
}
