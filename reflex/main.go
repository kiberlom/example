package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := struct {
		string
		int
	}{"dssa", 12}

	v := reflect.ValueOf(s)

	fmt.Printf("%#v\n", v)
	fmt.Println(v.Kind().String())
}
