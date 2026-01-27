package geometry_test

import (
	"testing"

	"github.com/yehezkiel1086/go-docs/unit-tests/go-test-geometry/geometry"
)

func handleTestError(t *testing.T, expected, actual float64) {
	t.Errorf("Expected %f, but got %f", expected, actual)
}

func TestArea(t *testing.T) {
	cube := &geometry.Cube{Side: 2}
	expected := 24.0

	if cube.Area() != expected {
		handleTestError(t, expected, cube.Area())
	}
}

func TestCircumference(t *testing.T) {
	cube := &geometry.Cube{Side: 2}
	expected := 24.0

	if cube.Circumference() != expected {
		handleTestError(t, expected, cube.Circumference())
	}
}

func TestVolume(t *testing.T) {
	cube := &geometry.Cube{Side: 2}
	expected := 8.0

	if cube.Volume() != expected {
		handleTestError(t, expected, cube.Volume())
	}
}
