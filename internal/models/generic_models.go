package models

// Error Response
// swagger:response errorMessage
type _ struct {
	// in:body
	Body ErrorMessage
}

type ErrorMessage struct {
	// Error message string
	// Example: error message
	Message string `json:"message"`
}
