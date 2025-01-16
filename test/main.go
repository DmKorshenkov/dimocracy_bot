package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Создаем новый канал для передачи данных
	lines := make(chan string)

	// Запускаем горутину для чтения с командной строки
	go func() {
		fmt.Println("goroutine")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			// Отправляем каждую введенную строку в канал
			lines <- scanner.Text()
		}
		fmt.Println("close goroutine")
		//close(lines) // Закрываем канал, когда ввод завершен
	}()

	// Основной поток
	for line := range lines {
		// Обрабатываем каждую строку, полученную из канала
		fmt.Println("Вы ввели:", line)
		if !strings.Contains(line, " ") {
			switch line {
			case "exit":
				close(lines)
			case "pwd":
				dir, _ := os.Getwd()
				fmt.Println(dir)
			case "ls":
				dirs, _ := os.ReadDir("./")
				if len(dirs) == 0 {
					fmt.Println("empty")
					continue
				}
				for _, dir := range dirs {
					fmt.Println(dir.Name())
				}
			}
		}
		if strings.Contains(line, " ") {
			lineSl := strings.Split(line, " ")
			switch lineSl[0] {
			case "cd":
				if lineSl[1] == ".." {
					os.Chdir("../")
				} else {
					err := os.Chdir("./" + lineSl[1])
					if err != nil {
						log.Println(err.Error())
					}
				}
			case "mkdir":
				err := os.Mkdir(lineSl[1], 0755)
				if err != nil {
					log.Println(err.Error())
				}
			}
		}
	}
	dir, _ := os.Getwd()
	fmt.Println(dir)
	fmt.Println("end")
}
