package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(message string, err error) Response {
	return Response{
		Message: message,
		Error:   err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is required", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("%s is not valid", err.Field()))
		}
	}
	return Response{
		Message: "Validation failed",
		Error:   strings.Join(errMsgs, ", "),
	}
}
