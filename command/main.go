package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

// запуск программы и вывод результата
func fileStart() {

	// определяем команду
	cmd := exec.Command("app_z.exe")


	// Output запускает и считывает вывод (2 в 1)
	b, err := cmd.Output()
	if err != nil {
		fmt.Println("Ошибка 2 ", err)
		return
	}

	fmt.Printf("%s", b)

}

// запуск программы и вывод результата в потоке
func readOutCmd() {
	// создаем команду
	cmd := exec.Command("app_z.exe")

	// создаем паток чтения
	r, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// стартуем команду
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// читаем вывод по байтам
	buf := make([]byte, 100)
	for {
		n, err := r.Read(buf)
		fmt.Print(string(buf[:n]))
		// ловим завершение работы команды
		if err == io.EOF {
			break
		}
	}

	// ждем завершения
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("")
}


func main() {

}
