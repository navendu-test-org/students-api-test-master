package handlers

import (
	"log"
)

// Details is a http handler
type Students struct {
	l *log.Logger
}

type KeyStudent struct{}

func NewStudents(l *log.Logger) *Students {
	return &Students{l}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
