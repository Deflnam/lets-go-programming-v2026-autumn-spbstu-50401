package main

import (
	"fmt"

	"github.com/Deflnam/task-3/internal/parser"
)

func main() {
	fmt.Println("Запуск программы для task-3 (Работа с библиотеками)")

	files := []string{
		"data/1.xml",
		"data/2.xml",
		"data/3.xml",
		"data/4.xml",
		"data/5.xml",
	}

	for _, file := range files {
		parser.ParseFile(file)
	}

	fmt.Println("Программа успешно завершена.")
}
