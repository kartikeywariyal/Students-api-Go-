package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string
	Error  string
}

func WriteJsonResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
func GenralErrorResponse(e error) Response {
	return Response{
		Status: "Failed",
		Error:  e.Error(),
	}
}
func ValidationErrorResponse(e validator.ValidationErrors) Response {
	var errorMsg []string
	for _, err := range e {
		errorMsg = append(errorMsg, err.Field()+" is "+err.Tag())
		return Response{
			Status: "Failed",
			Error:  err.Error(),
		}
	}
	return Response{
		Status: "Failed",
		Error:  "Validation errors: " + fmt.Sprint(errorMsg, ','),
	}
}
