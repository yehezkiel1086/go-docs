package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	DateBirth time.Time
}

func (p *Person) GetAge() int {
	return time.Now().Year() - p.DateBirth.Year()
}

type Student struct {
	Person
	Batch uint
	GPA float64
}

func (s *Student) GetPredicate() string {
	switch {
	case s.GPA >= 3.5:
		return "A"
	case s.GPA >= 3.0:
		return "AB"
	case s.GPA >= 2.5:
		return "B"
	case s.GPA >= 2.0:
		return "C"
	default:
		return "F"
	}
}

func main() {
	john := Student{
		Person: Person{
			Name: "John",
			DateBirth: time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Batch: 2020,
		GPA: 3.4,
	}

	fmt.Println(john.GetAge(), john.GetPredicate())
}