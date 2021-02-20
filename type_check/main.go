package main

import "fmt"

// проверка преобразоватие типа

func main() {

	var i interface{} = 2.0

	p, ok := i.(float64)

	if !ok {
		fmt.Println("Не правильно указан тип")
		return
	}

	fmt.Printf("%T  %v", p, p)
}
