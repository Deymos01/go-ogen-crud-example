package internal

import "fmt"

type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Car with ID %d not found", e.ID)
}
