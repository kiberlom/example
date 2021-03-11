package main

func Silo() []int {
	var m []int
	for i := 0; i < 1000000; i++ {
		m = append(m, i)
	}

	return m

}

func Filo() []int {
	var m []int
	for i := 0; i < 100; i++ {
		m = append(m, i)
	}

	return m

}

func main() {

}
