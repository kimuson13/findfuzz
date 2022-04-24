package a_test

import (
	"fmt"
	"testing"
)

func f() {
	a := 0
	fmt.Println(a)
}

func FuzzF(f *testing.F) { // want "Fuzz test here"
	fmt.Println("fuzz")
}

func FuzzNG(t *testing.T) {
	fmt.Println("no fuzz")
}
