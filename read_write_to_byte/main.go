package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	//
	f2, err := os.Create("main4_t.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f2.Close()

	fi, _ := f.Stat()
	fmt.Println(fi.Size())

	//fmt.Printf("%#v", f)

	buf := make([]byte, 15)

	countW := 0

	for {

		n, err := f.Read(buf)
		if err == io.EOF {
			fmt.Println(err)
			break
		}

		fmt.Println(string(buf[:n]))

		n, err = f2.Write(buf[:n])

		countW += n
		if err != nil {
			fmt.Println(err)
			break
		}

	}

	fmt.Println(countW)

}

