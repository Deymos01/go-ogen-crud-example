package api

import (
	"context"
	"github.com/Deymos01/go-ogen-crud-example/internal/oas"
	"sort"
	"sync"
)

var _ oas.Handler = (*CarHandler)(nil)

type CarHandler struct {
	mu   sync.Mutex
	data map[int]oas.Car
	id   int
}

func NewCarHandler() *CarHandler {
	return &CarHandler{
		data: make(map[int]oas.Car),
		id:   0,
	}
}

func (c *CarHandler) AddCar(ctx context.Context, req *oas.NewCar) (*oas.Car, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	car := oas.Car{
		ID:           c.id,
		Manufacturer: req.Manufacturer,
		Model:        req.Model,
		Year:         req.Year,
		Color:        req.Color,
	}

	c.data[c.id] = car
	c.id++
	return &car, nil
}

func (c *CarHandler) DeleteCarById(ctx context.Context, params oas.DeleteCarByIdParams) (oas.DeleteCarByIdRes, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[params.ID]; !ok {
		return ErrNotFound(params.ID), nil
	}
	delete(c.data, params.ID)
	return &oas.DeleteCarByIdNoContent{}, nil
}

func (c *CarHandler) GetCarById(ctx context.Context, params oas.GetCarByIdParams) (oas.GetCarByIdRes, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	car, ok := c.data[params.ID]
	if !ok {
		return ErrNotFound(params.ID), nil
	}
	return &car, nil
}

func (c *CarHandler) ListCars(ctx context.Context) ([]oas.Car, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cars := make([]oas.Car, 0, len(c.data))
	for _, car := range c.data {
		cars = append(cars, car)
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].ID < cars[j].ID
	})
	return cars, nil
}

func (c *CarHandler) UpdateCarById(ctx context.Context, req *oas.UpdateCar, params oas.UpdateCarByIdParams) (oas.UpdateCarByIdRes, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	car, ok := c.data[params.ID]
	if !ok {
		return ErrNotFound(params.ID), nil
	}

	if req.Model.IsSet() {
		car.Model = req.Model.Value
	}
	if req.Year.IsSet() {
		car.Year = req.Year.Value
	}
	if req.Manufacturer.IsSet() {
		car.Manufacturer = req.Manufacturer.Value
	}
	if req.Color.IsSet() {
		car.Color = req.Color.Value
	}

	c.data[params.ID] = car
	return &car, nil
}
