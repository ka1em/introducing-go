package main

import (
	"fmt"
	"hash/crc32"
	"io"
	//"io/ioutil"
	"os"
)

func getHash(filename string) (uint32, error) {
	//open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, nil
	}

	//close opend files
	defer f.Close()

	//create a hasher
	h := crc32.NewIEEE()

	//copy the file into the hasher
	_, err = io.Copy(h, f)

	if err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("file.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	h2, err := getHash("file2.go")
	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}
