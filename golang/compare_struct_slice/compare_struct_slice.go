package main

import (
	"fmt"
	"reflect"
)

// https://paulownia.hatenablog.com/entry/2017/08/28/221515

type Sample struct {
	t bool
	d int
	s string
}

func main() {
	a := Sample{t: true, d: 1, s: "test"}
	b := Sample{t: true, d: 1, s: "test"}
	c := Sample{t: false, d: 1, s: "test"}

	fmt.Printf("%t\n", reflect.DeepEqual(&a, &b)) // true
	fmt.Printf("%t\n", reflect.DeepEqual(&a, &c)) // false
}
