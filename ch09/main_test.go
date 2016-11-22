package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	v := Add(3, 3)
	if v != 6 {
		t.Error("add error")
	}
}
