package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var number int32 = 10
	var isMaried bool = false

	// conversion operator
	var bigNum = int64(number) // convert int32 to int64
	var floatNum = float32(number) // convert int32 to float32

	// Sprintf - converts to string
	var numStr = fmt.Sprintf("%d", number) // convert int32 to string
	var isMariedStr = fmt.Sprintf("%t", isMaried) // convert bool to string
	var Pi = fmt.Sprintf("%f", math.Pi) // convert float to string

	// strconv.Itoa - int to string
	var str = strconv.Itoa(10)

	// strconf.Atoi - string to int
	var num, err = strconv.Atoi("10")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%T %v\n", bigNum, bigNum)
	fmt.Printf("%T %v\n", floatNum, floatNum)
	fmt.Printf("%T %v\n", numStr, numStr)
	fmt.Printf("%T %v\n", isMariedStr, isMariedStr)
	fmt.Printf("%T %v\n", Pi, Pi)
	fmt.Printf("%T %v\n", str, str)
	fmt.Printf("%T %v\n", num, num)
}
