package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

type Student struct {
	Person Person // embedded struct
	Major  string
	GPA    float32
}

func (s Student) String() string {
	return fmt.Sprintf("Name: %s\nAge: %d\nMajor: %s\nGPA: %.2f", s.Person.Name, s.Person.Age, s.Major, s.GPA)
}

func (s *Student) UpdateGPA(gpa float32) {
	s.GPA = gpa
}

func main() {
	s := &Student{
		Person: Person{
			Name: "Yehezkiel",
			Age:  23,
		},
		Major:  "Computer Science",
		GPA:    3.4,
	}
	s.UpdateGPA(3.34)
	fmt.Println(s)
}
