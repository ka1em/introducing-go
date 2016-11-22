package main

import (
	"fmt"
)

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("WRON")
	//str := recover() //this will nerver happen
	//fmt.Println(str)
}
