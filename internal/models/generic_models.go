package models

// ErrorMessage is a generic error message returned by a server
type ErrorMessage struct {
	Message string `json:"message"`
}
