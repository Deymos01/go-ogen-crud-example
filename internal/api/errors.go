package api

import (
	"fmt"
)

type CarNotFoundError struct {
	ID int
}

func (e *CarNotFoundError) Error() string {
	return fmt.Sprintf("car with ID %d not found", e.ID)
}
