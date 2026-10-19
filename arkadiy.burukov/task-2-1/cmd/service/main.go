package main

import "fmt"

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	for range departments {
		var employees int

		_, err = fmt.Scan(&employees)
		if err != nil {
			fmt.Println("Invalid input", err)

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
				fmt.Println("Invalid input", err)

				return
			}

			if operator == ">=" && temp > minTemp {
				minTemp = temp
			}

			if operator == "<=" && temp < maxTemp {
				maxTemp = temp
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
