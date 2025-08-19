package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var num int32 = 10
	var isMarried bool = false

	// convert to string
	numStr := fmt.Sprintf("%d", num)
	isMarriedStr := fmt.Sprintf("%t", isMarried)

	fmt.Printf("%s %s\n", reflect.TypeOf(numStr), reflect.TypeOf(isMarriedStr))

	// int to string
	age := 21
	strAge := strconv.Itoa(age)
	ageAgain, _ := strconv.Atoi(strAge)
	fmt.Printf("Type of age: %s\nType of strAge: %s\nType of strAgain: %s\n", reflect.TypeOf(age), reflect.TypeOf(strAge), reflect.TypeOf(ageAgain))
}
