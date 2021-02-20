package main

import (
	"fmt"
	"io"
	"strings"
)

func readBuf() {
	r := strings.NewReader("Hello World")

	buf := make([]byte, 4)

	for {

		n, err := r.Read(buf)
		if err == io.EOF {
			fmt.Println(err)
			break
		}

		fmt.Println(string(buf[:n]), "  ", n)
	}
}

func main() {

	s := "Привtет"

	rs := []rune{'H', 'e', 'l', 'l', 'o', ' ', '世', '界'}

	fmt.Println(rs)

	for _, r := range s {
		fmt.Println([]byte(string(r)))
	}

}
