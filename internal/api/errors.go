package api

import (
	"fmt"
	"github.com/Deymos01/go-ogen-crud-example/internal/oas"
	"net/http"
)

func ErrNotFound(id int) *oas.Error {
	return &oas.Error{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("Car with ID %d not found", id),
	}
}
