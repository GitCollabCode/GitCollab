package models

// Error Response
// swagger:response errorResponse
type _ struct {
	// in:body
	Body ErrorMessage
}

type ErrorMessage struct {
	// Error message string
	// Example: error message
	Message string `json:"message"`
}

// Message Response
// swagger:response messageResponse
type _ struct {
	// in:body
	Body ErrorMessage
}

type Message struct {
	// message string
	// Example: resource created
	Message string `json:"message"`
}
