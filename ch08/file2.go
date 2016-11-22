package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("../ch01/env.sh")
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
