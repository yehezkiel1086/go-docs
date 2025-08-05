package methods

import (
	"fmt"
	"os"
)

type File struct {
	Filename string
}

func (f *File) CreateFile() {
	if _, err := os.Stat(f.Filename); os.IsNotExist(err) {
		// create new file
		file, err := os.Create(f.Filename);
		if err != nil {
			panic(err)
		}

		fmt.Printf("File %v created!", file.Name())
		defer file.Close()
	} else {
		fmt.Println("File already exists")
	}
}

func (f *File) WriteContent(content string) {
	if err := os.WriteFile(f.Filename, []byte(content), 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Content written to %v successfully!\n", f.Filename)
}

func (f *File) ReadContent() (string) {
	content, err := os.ReadFile(f.Filename)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func (f *File) RemoveFile() {
	if err := os.Remove(f.Filename); err != nil {
		panic(err)
	}
	fmt.Println("File successfully removed.")
}
