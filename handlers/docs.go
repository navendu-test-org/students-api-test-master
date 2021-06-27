// Package classification of Student API
//
// Documentation for Student API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta

package handlers

import "github.com/navendu-pottekkat/students-api-test/data"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of students
// swagger:response studentsResponse
type studentsResponseWrapper struct {
	// All current students
	// in: body
	Body []data.Student
}

// Data structure representing a single student
// swagger:response studentResponse
type studentResponseWrapper struct {
	// Newly created student
	// in: body
	Body data.Student
}
