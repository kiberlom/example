package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// анализ введенных данных
func main() {

	fmt.Print("Введите число: ")

	var x string

	for {

		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		if s.Err() != nil {
			fmt.Println("Еще разок")
			continue
		}

		x = s.Text()

		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			fmt.Print("Введите число: ")
			continue
		}

		fmt.Printf("число: %f   типа: %T \n", f, f)
		break

	}
}
