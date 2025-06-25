package main

import "io-test/filecli"

// consts
const FILENAME = "hello.txt"
const CONTENT = "Hello File! This is Golang writing."

func main() {
	filecli.CreateEmptyFile(FILENAME)
	filecli.WriteContent(FILENAME, CONTENT)
	filecli.ReadFile(FILENAME)
}
