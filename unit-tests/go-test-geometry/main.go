package main

import (
	"fmt"

	"github.com/yehezkiel1086/go-docs/unit-tests/go-test-geometry/geometry"
)

func main() {
	cube := &geometry.Cube{Side: 2}

	fmt.Println(cube.Area(), cube.Circumference(), cube.Volume())
}
