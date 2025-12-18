package main

import (
	"testing"
	"time"
)

func TestGetAge(t *testing.T) {
	p := &Person{
		Name: "John Doe",
		DateBirth: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	expected := 45

	if p.GetAge() != expected {
		t.Errorf("Expected %d, got %d", expected, p.GetAge())
	}
}

func TestGetPredicate(t *testing.T) {
	st := Student{
		Person: Person{
			Name: "Jane Doe",
			DateBirth: time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Batch: 2014,
		GPA: 3.6,
	}
	expected := "A"

	if expected != st.GetPredicate() {
		t.Errorf("Expected %s, got %s", expected, st.GetPredicate())
	}
}