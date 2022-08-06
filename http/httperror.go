package http

import (
	"encoding/json"
	"net/http"
	"strings"
)

type err struct {
	Error   string   `json:"error"`
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Cause   []string `json:"cause"`
}

func Render(w http.ResponseWriter, statusCode int, message string, causes ...string) {
	parsedErr := strings.ToLower(strings.ReplaceAll(http.StatusText(statusCode), " ", "_"))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	err := err{
		Error:   parsedErr,
		Status:  statusCode,
		Message: message,
		Cause:   causes,
	}

	json.NewEncoder(w).Encode(err) // nolint:errcheck
}

func RenderFromApiError(w http.ResponseWriter, apiError ApiError) {
	Render(w, apiError.Status(), apiError.Message(), apiError.Cause().ToString())
}

func RenderInternalServerError(w http.ResponseWriter, err error) {
	Render(
		w,
		http.StatusInternalServerError,
		"There's an error while processing your request. Please check the cause to identify the problem",
		err.Error(),
	)
}
