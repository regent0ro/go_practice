/*
	Print Fibonacci numbers
	Use iteration
*/

package main

import "fmt"

func main() {
	var n int

	fmt.Print("input N: ")
	fmt.Scan(&n)

	if n <= 0 {
		return
	}

	p, q := 0, 1 // short variable declaration
	for i := 0; i <= n; i++ {
		p, q = q, p+q
	}

	fmt.Print("Fibonacci ", n, ":", q)
}
