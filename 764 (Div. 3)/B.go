package main

import (
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err == nil {
		os.Stdin = file
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

}
