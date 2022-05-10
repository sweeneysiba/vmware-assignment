package forms

import (
	"testing"
)

type FormTest struct {
	Tag   string
	Error string
}

var sort = []FormTest{
	{"required", "please specify the required field"},
	{"tempTag", "Something went wrong, please try again later"},
}
var limit = []FormTest{
	{"required", "please specify the required field"},
	{"tempTag", "Something went wrong, please try again later"},
}

func TestForm(t *testing.T) {
	assignmentForm := new(AssignmentForm)
	for _, val := range sort {
		if err := assignmentForm.Sort(val.Tag, val.Error); err != val.Error {
			t.Fatalf("Expected error %s but got %s", val.Error, err)
		}
	}
	for _, val := range limit {
		if err := assignmentForm.Limit(val.Tag, val.Error); err != val.Error {
			t.Fatalf("Expected error %s but got %s", val.Error, err)
		}
	}
}
