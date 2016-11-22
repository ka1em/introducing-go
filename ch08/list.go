package main

import (
	"container/list"
	"fmt"
)

var x list.List

type stu struct {
	Name string
	Num  int
}

func main() {

	stu1 := stu{
		Name: "wang",
		Num:  1,
	}

	sut2 := stu{
		Name: "zhang",
		Num:  2,
	}

	x.PushBack(stu1)
	x.PushBack(sut2)

	for e := x.Front(); e != nil; e = e.Next() {
		tmp := e.Value.(stu)
		fmt.Println(tmp.Name, tmp.Num)
	}
}
