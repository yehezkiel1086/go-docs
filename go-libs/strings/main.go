package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToLower("YEHEZKIEL"))
	fmt.Println(strings.ToUpper("yehezkiel"))

	fmt.Println(strings.Replace("Yehezkiel", "e", "a", 1)) // replace 1 e characters with a
	fmt.Println(strings.Replace("Yehezkiel", "e", "a", 2)) // replace 2 e characters with a
	fmt.Println(strings.Replace("Yehezkiel", "e", "a", -1)) // replace all e characters with a

	arr := strings.Split("Hello Go!", " ")
	fmt.Println(arr)
	fmt.Println(strings.Join(arr, "-"))

	fmt.Println(strings.Contains("Yehezkiel Wiradhika", "kiel"))
	fmt.Println(strings.HasPrefix("Yehezkiel", "Yehez"))
	fmt.Println(strings.HasSuffix("Wiradhika", "dhika"))

	fmt.Println(strings.Count("Yehezkiel", "e"))
	fmt.Println(strings.Index("Yehezkiel", "k"))

	fmt.Println(strings.Repeat("Ho", 3))
}
