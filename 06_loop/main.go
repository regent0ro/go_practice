package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	for index, number := range numbers {
		fmt.Println(index, number)
	}

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(total)
}
