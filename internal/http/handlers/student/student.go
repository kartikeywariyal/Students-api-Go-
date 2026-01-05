package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kartikeywariyal/students-api-Go-/internal/storage"
	"github.com/kartikeywariyal/students-api-Go-/internal/utils/response"

	"github.com/kartikeywariyal/students-api-Go-/internal/types"
)

func New(db storage.Storage) http.HandlerFunc {
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
		_, err = db.CreateStudent(student.Name, student.Age, student.Email)
		if err != nil {
			response.WriteJsonResponse(w, http.StatusInternalServerError, response.GenralErrorResponse(err))
			return
		}
		response.WriteJsonResponse(w, http.StatusOK, map[string]string{"message": "Student created successfully"})
	})
}
func GetStudent(db storage.Storage) http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJsonResponse(w, http.StatusBadRequest, response.GenralErrorResponse(err))
			return
		}

		student, err := db.GetStudent(idInt)
		if err != nil {
			response.WriteJsonResponse(w, http.StatusInternalServerError, response.GenralErrorResponse(err))
			return
		}
		response.WriteJsonResponse(w, http.StatusOK, student)

	})
}
