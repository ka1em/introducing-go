package main

import (
	"fmt"
)

func main() {
	x := make(map[string]string)
	y := map[string]map[string]string{
		"A": map[string]string{
			"AA": "AAA",
		},
	}

	x["a"] = "aa"
	x["b"] = "bb"

	fmt.Println(x["a"], x["b"])

	if yy, ok := y["A"]; ok {
		fmt.Println(yy["AA"])
	}
}
