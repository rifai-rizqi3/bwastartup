package helper

import "github.com/go-playground/validator/v10"

// Response is a struct representing the standardized API response format.
type Response struct {
	Meta Meta        `json:"meta"` // Meta information about the response.
	Data interface{} `json:"data"` // Data payload of the response.
}

// Meta is a struct representing metadata about the API response.
type Meta struct {
	Message string `json:"message"` // A message describing the result of the operation.
	Code    int    `json:"code"`    // An integer code indicating the status or result of the operation.
	Status  string `json:"status"`  // A string indicating the status of the operation (e.g., "success" or "error").
}

// APIResponse is a function that creates and returns a standardized API response.
func APIResponse(message string, code int, status string, data interface{}) Response {
	// Create Meta structure with provided parameters.
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	// Create Response structure with Meta and Data.
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

// FormatValidationError is a function that formats validation errors into a slice of strings.
func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
