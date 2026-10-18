package main

import "fmt"

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	for range departments {
		var employees int

		_, err = fmt.Scan(&employees)
		if err != nil {
			fmt.Println("Invalid input")

			return
		}

		minTemp := 15
		maxTemp := 30

		for range employees {
			var (
				operator string
				temp     int
			)

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
