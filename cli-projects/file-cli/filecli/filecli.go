package filecli

import (
	"fmt"
	"io"
	"os"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)

	return os.IsExist(err)
}

func OpenFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)

	if err != nil {
		panic(err)
	}

	return file
}

func CreateEmptyFile(filename string) {
	exists := FileExists(filename)
	
	if !exists {
		file, err := os.Create(filename)

		if err != nil {
			panic(err)
		}

		defer file.Close()
		
		fmt.Printf("File %v created!\n", file.Name())
	} else {
		fmt.Println("File exists!")
	}
}

func WriteContent(filename string, content string) {
	file := OpenFile(filename)

	defer file.Close()

	n, err := file.WriteString(content)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v bytes of content written!\n", n)
}

func ReadFile(filename string) {
	file := OpenFile(filename)

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
