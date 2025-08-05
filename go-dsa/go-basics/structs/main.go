package main

import "fmt"

type Person struct {
	Name string
	Age uint8
}

func (p *Person)changeName(name string) {
	p.Name = name
}

func (p Person)sayHello() {
	fmt.Printf("Hello, %v\n", p.Name)
}

func (p Person)String() (string) {
	return fmt.Sprintf("The name is %v, age is %v", p.Name, p.Age)
}

func main() {
	p := Person{
		Name: "Benjamin",
		Age: 21,
	}
	p.sayHello()
	p.changeName("Ben")
	fmt.Println(p)
}