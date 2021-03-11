package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := os.Getenv("PATH")
	fmt.Println(s)
	words := strings.Split(s, ";")
	for _, w := range words {

		ns := strings.Trim(w, " ")
		if ns != "" {
			fmt.Println(w)
		}

	}
}
