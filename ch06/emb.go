package main

import (
	"fmt"
)

type Persion struct {
	Name string
}

func (p *Persion) Talk() {
	fmt.Println("My name is ", p.Name)
}

type Android struct {
	Persion Persion
	Model   string
}

type Ios struct {
	Persion
	Model string
}

func main() {
	a := new(Android)
	a.Persion.Name = "wkm"
	a.Persion.Talk()
	b := new(Ios)

	b.Persion.Name = "sds"
	b.Persion.Talk()
}
