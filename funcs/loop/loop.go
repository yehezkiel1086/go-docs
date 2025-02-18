package funcs

import "fmt"

func ForLoop() {
	fmt.Println("For Loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Index: %v\n", i)
	}
	fmt.Println()
}