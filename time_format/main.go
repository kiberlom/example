package main

import (
	"fmt"
	"time"
)

func main() {
	d1 := "2021-01-14T15:12:51.6650085Z"
	layout := "2006-01-02T15:04:05.0000000Z"

	t, err := time.Parse(layout, d1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	loc, _ := time.LoadLocation("Europe/Istanbul")
	t = t.In(loc)
	n := t.Format("2006-01-02T15:04:05.0000000-07:00")
	fmt.Println(n)

	//fmt.Println("############################################")
	//now := time.Now()
	//n:=now.Format("2006-01-02T15:04:05.0000000-07:00")
	//fmt.Println(n)

}
