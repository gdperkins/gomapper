package gomap

import (
	"testing"
)

type employeeSource struct {
	FirstName string
	LastName  string
	Salary    int
}

type employeeDestination struct {
	FirstName string
	LastName  string
}

func TestCanIgnore(t *testing.T) {
	gm := New(Options{
		IgnoreFields: []string{"FirstName"},
	})

	source := employeeSource{"John", "Doe", 1000}
	destination := employeeDestination{}

	gm.Map(source, &destination)

	if destination.FirstName != "" {
		t.Errorf("Test failed, LastName should be empty.")
	}
}
