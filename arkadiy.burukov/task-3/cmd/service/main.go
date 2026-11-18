package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Deflnam/task-3/internal/parser"
)

func main() {
	files := []string{
		"data/3.xml",
		"data/4.xml",
		"data/1.xml",
		"data/5.xml",
		"data/2.xml",
	}

	fmt.Printf("%-12s | %-8s | %s\n", "Дата", "Номинал", "Курс USD (руб.)")
	fmt.Println(strings.Repeat("-", 40))

	for _, file := range files {
		vc, err := parser.ParseFile(file)
		if err != nil {
			log.Printf("Ошибка при чтении %s: %v", file, err)
			continue
		}

		for _, v := range vc.Valutes {
			if v.CharCode == "USD" {
				raw := strings.ReplaceAll(v.Value, ",", ".")
				rate, err := strconv.ParseFloat(raw, 64)
				if err != nil {
					log.Printf("Ошибка парсинга курса для %s: %v", vc.Date, err)
					continue
				}

				fmt.Printf("%-12s | %-8d | %.4f\n", vc.Date, v.Nominal, rate)
				break
			}
		}
	}
}
