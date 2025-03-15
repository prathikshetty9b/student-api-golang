package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/prathikshetty9b/students-api/pkg/storage"
	"github.com/prathikshetty9b/students-api/pkg/types"
	"github.com/prathikshetty9b/students-api/pkg/utils/response"
)

// New returns an HTTP handler for creating a new student
func New(db storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// Ensure the request method is POST
		if req.Method != http.MethodPost {
			slog.Error("Invalid request method")
			response.WriteJSON(res, http.StatusMethodNotAllowed, map[string]string{"error": "Invalid request method"})
			return
		}

		slog.Info("Creating a new student")

		// Decode the request body into a Student struct
		var student types.Student
		err := json.NewDecoder(req.Body).Decode(&student)
		if err != nil {
			if errors.Is(err, io.EOF) {
				slog.Error("Request body is empty")
				response.WriteJSON(res, http.StatusBadRequest, response.GeneralError("Request body is empty", err))
			} else {
				slog.Error("Failed to decode request body", slog.String("error", err.Error()))
				response.WriteJSON(res, http.StatusBadRequest, response.GeneralError("Failed to decode request body", err))
			}
			return
		}
		// request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			slog.Error("Validation failed", slog.String("error", validateErrs.Error()))
			response.WriteJSON(res, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}
		lastId, err := db.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			slog.Error("Failed to create student", slog.String("error", err.Error()))
			response.WriteJSON(res, http.StatusInternalServerError, response.GeneralError("Failed to create student", err))
			return
		}

		// Respond with a success message
		slog.Info("Student created successfully", slog.Int64("id", lastId))
		response.WriteJSON(res, http.StatusCreated, map[string]any{
			"message": "Student created successfully",
			"id":      lastId,
		})
	}
}

func GetById(db storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// Ensure the request method is GET
		if req.Method != http.MethodGet {
			slog.Error("Invalid request method")
			response.WriteJSON(res, http.StatusMethodNotAllowed, response.GeneralError("Invalid request method", nil))
			return
		}

		slog.Info("Fetching student by id")
		// Get the student ID from the URL path
		id := req.PathValue("id")
		if id == "" {
			slog.Error("Student ID is required")
			response.WriteJSON(res, http.StatusBadRequest, response.GeneralError("Student ID is required", nil))
			return
		}

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("Invalid student ID", slog.String("error", err.Error()))
			response.WriteJSON(res, http.StatusBadRequest, response.GeneralError("Invalid student ID", err))
			return
		}

		student, err := db.GetStudentByID(intId)
		if err != nil {
			slog.Error("Failed to fetch student", slog.String("error", err.Error()))
			response.WriteJSON(res, http.StatusInternalServerError, response.GeneralError("Failed to fetch student", err))
			return
		}

		// Respond with the student details
		slog.Info("Student fetched successfully", slog.String("id", id))
		response.WriteJSON(res, http.StatusOK, student)
	}
}

func GetList(db storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// Ensure the request method is GET
		if req.Method != http.MethodGet {
			slog.Error("Invalid request method")
			response.WriteJSON(res, http.StatusMethodNotAllowed, response.GeneralError("Invalid request method", nil))
			return
		}

		slog.Info("Fetching all students")
		// Fetch all students from the database
		students, err := db.GetStudents()
		if err != nil {
			slog.Error("Failed to fetch students", slog.String("error", err.Error()))
			response.WriteJSON(res, http.StatusInternalServerError, response.GeneralError("Failed to fetch students", err))
			return
		}

		// Respond with the list of students
		slog.Info("Students fetched successfully")
		response.WriteJSON(res, http.StatusOK, students)
	}
}
