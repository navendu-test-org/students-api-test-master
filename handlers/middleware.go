package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/navendu-pottekkat/students-api-test/data"
)

func (s Students) MiddlewareValidateStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		stud := data.Student{}

		err := stud.FromJSON(r.Body)

		if err != nil {
			s.l.Println("[Error] deserialising student", err)
			http.Error(
				rw,
				"Error reading student",
				http.StatusBadRequest,
			)
			return
		}

		// Validate the student
		err = stud.Validate()
		if err != nil {
			s.l.Println("[ERROR] validating student", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating student: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// Add the student to the context
		ctx := context.WithValue(r.Context(), KeyStudent{}, stud)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(rw, r)
	})
}
