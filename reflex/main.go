package main

import "fmt"

func d(y int) {
	y++
	fmt.Println(y)
}

func main() {
	x := 1
	d(x)
	fmt.Println(x)
}
