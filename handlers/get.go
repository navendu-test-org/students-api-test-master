package handlers

import (
	"net/http"

	"github.com/navendu-pottekkat/students-api-test/data"
)

// swagger:route GET / students listStudents
// Return a list of students from the database
// responses:
//	200: studentsResponse

// GetStudents handles all the GET requests and returns a list of students
func (s *Students) GetStudents(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle GET Students")

	ls := data.GetStudents()

	// Serialize to JSON
	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
