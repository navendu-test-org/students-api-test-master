package data

import (
	"encoding/json"
	"io"
	"regexp"

	"github.com/go-playground/validator"
)

type Student struct {
	Name   string `json:"name" validate:"required"`
	ID     string `json:"id" validate:"required,id"`
	Age    int    `json:"age"`
	Course string `json:"course"`
}

type Students []*Student

var StudentsList = Students{
	&Student{
		Name:   "Jon Snow",
		ID:     "abc-fhi-jke",
		Age:    15,
		Course: "Sparring",
	},
	&Student{
		Name:   "Robb Stark",
		ID:     "abc-gkw-iml",
		Age:    15,
		Course: "Sparring",
	},
	&Student{
		Name:   "Bran Stark",
		ID:     "abc-gkw-klo",
		Age:    7,
		Course: "Climbing",
	},
	&Student{
		Name:   "Rickon Stark",
		ID:     "abc-gkw-ibh",
		Age:    3,
		Course: "Playing",
	},
	&Student{
		Name:   "Sansa Stark",
		ID:     "abc-gkw-qow",
		Age:    11,
		Course: "Gentle Arts",
	},
	&Student{
		Name:   "Arya Stark",
		ID:     "abc-gkw-few",
		Age:    9,
		Course: "Gentle Arts",
	},
}

func (s *Students) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

func (s *Student) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(s)
}

func (s *Student) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("id", validateID)
	return validate.Struct(s)
}

func validateID(fl validator.FieldLevel) bool {
	// ID is of the form abc-def-jhi
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

func GetStudents() Students {
	return StudentsList
}

func AddStudent(s *Student) {
	StudentsList = append(StudentsList, s)
}
