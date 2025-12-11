package main

import "testing"

func TestArea(t *testing.T) { // fungsi ini digunakan untuk mengetes method Area di struct 'Cube'
    var cube = Cube{Side: 2} // menyiapkan data untuk diuji
    var expected float64 = 24.0 // nilai yang diharapkan

    if cube.Area() != expected { // jika hasil tidak sesuai dengan yang diharapkan, maka test dianggap gagal
        t.Errorf("Expected %f, but got %f", expected, cube.Area())
    }
}

func TestCircumference(t *testing.T) {
	cube := Cube{Side: 2}
	var expected float64 = 24

	if cube.Circumference() != expected {
		t.Errorf("Expected %f, but got %f", expected, cube.Area())
	}
}

func TestVolume(t *testing.T) {
	cube := Cube{Side: 2}
	var expected float64 = 8

	if cube.Volume() != expected {
		t.Errorf("expected %f, but got %f", expected, cube.Volume())
	}
}
