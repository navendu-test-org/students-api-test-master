package handlers

import (
	"net/http"

	"github.com/navendu-pottekkat/students-api-test-master/data"
)

// swagger:route POST / students createStudent
// Create a new product
//
// responses:
//	200: studentResponse
//  422: errorValidation
//  501: errorResponse

// AddStudent handles POST requests to add new students
func (s *Students) AddStudent(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle POST Student")
	stud := r.Context().Value(KeyStudent{}).(data.Student)
	data.AddStudent(&stud)
}
