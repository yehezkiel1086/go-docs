package main

import (
	"fmt"
	"os"
)

func createFile(filename string) {
	// check if file exists
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		// create new file
		file, err := os.Create("lorem.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		fmt.Printf("File %s created successfully!\n", file.Name())
	} else {
		fmt.Println("File already exists.")
	}
}

func writeContent(filename string, content string) {
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		panic(err)
	}
}

func readContent(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func deleteFile(filename string) {
	if err := os.Remove(filename); err != nil {
		panic(err)
	}

	fmt.Println("File deleted successfully.")
}

func main() {
	FILENAME := "lorem.txt"
	// createFile(FILENAME)
	// writeContent(FILENAME, libs.RandStr(1000))
	// readContent(FILENAME)

	deleteFile(FILENAME)
}
