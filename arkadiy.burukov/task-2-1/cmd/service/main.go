package main

import "fmt"

func main() {
	var departments int
	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	for i := 0; i < departments; i++ {
		var employees int
		_, err = fmt.Scan(&employees)
		if err != nil {
			fmt.Println("Invalid input")

			return
		}

		minTemp := 15
		maxTemp := 30

		for j := 0; j < employees; j++ {
			var operator string
			var temp int

			_, err = fmt.Scan(&operator, &temp)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			if operator == ">=" {
				if temp > minTemp {
					minTemp = temp
				}
			}

			if operator == "<=" {
				if temp < maxTemp {
					maxTemp = temp
				}
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
