package main

import (
	"fmt"
	"os"
)

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return
	}

	for range departments {
		var employees int

		_, err := fmt.Scan(&employees)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)

			return
		}

		minTemp := 15
		maxTemp := 30

		var clientComparator string

		var clientWish int
		for range employees {
			_, err := fmt.Scan(&clientComparator, &clientWish)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)

				return
			}

			switch clientComparator {
			case ">=":
				minTemp = max(minTemp, clientWish)
			case "<=":
				maxTemp = min(maxTemp, clientWish)
			default:
				fmt.Fprintln(os.Stderr, "Incorrect comparator sign")
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
