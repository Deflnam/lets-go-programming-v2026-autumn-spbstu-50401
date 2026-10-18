package main

import "fmt"

func main() {
	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	for i := 0; i < n; i++ {
		var k int
		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		low := 15
		high := 30

		for j := 0; j < k; j++ {
			var op string
			var t int
			_, err = fmt.Scan(&op, &t)
			if err != nil {
				fmt.Println("Invalid input")
				return
			}

			if op == ">=" {
				if t > low {
					low = t
				}
			}
			if op == "<=" {
				if t < high {
					high = t
				}
			}
			if low > high {
				fmt.Println(-1)
			} else {
				fmt.Println(low)
			}

		}

	}
}
