package main

import (
	"os"
)

func main() {
	file, err := os.Create("test")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("testttt")
}
