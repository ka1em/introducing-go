package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func aa(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}
func main() {
	filepath.Walk("..", aa)
}
