package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kartikeywariyal/students-api-Go-/internal/utils/response"

	"github.com/kartikeywariyal/students-api-Go-/internal/types"
)

func New() http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJsonResponse(w, http.StatusBadRequest, response.GenralErrorResponse(err))
			return
		}
		if err != nil {
			response.WriteJsonResponse(w, http.StatusBadGateway, response.GenralErrorResponse(err))
			return
		}

		// request validation

		var validate = validator.New()

		if err := validate.Struct(student); err != nil {
			response.WriteJsonResponse(w, http.StatusBadRequest, response.ValidationErrorResponse(err.(validator.ValidationErrors)))
			return
		}

		response.WriteJsonResponse(w, http.StatusOK, map[string]string{"message": "Student created successfully"})
	})
}
