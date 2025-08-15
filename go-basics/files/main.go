package main

import (
	"fmt"
	"go-files/methods"
)

func main() {
	// create new file
	file := methods.File{
		Filename: "lorem.txt",
	}

	file.CreateFile()
	file.WriteContent("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
	fmt.Println(file.ReadContent())
	file.RemoveFile()
}
