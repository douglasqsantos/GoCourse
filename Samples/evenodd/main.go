package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// var numbers []int

	// for i := 0; i <= 10; i++ {
	// 	numbers = append(numbers, i)
	// }

	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Printf("%v is even\n", n)
			// fmt.Println(n, "is even")
		} else {
			fmt.Printf("%v is odd\n", n)
			// fmt.Println(n, "is odd")
		}

	}
}
