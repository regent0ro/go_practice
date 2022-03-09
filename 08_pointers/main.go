package main

import "fmt"

func main() {
	a := 2
	b := a //copy of value

	a = 10
	fmt.Println(&a, &b) //address of a, address of b
	fmt.Println(a, b)   // 10 2

	c := 2
	d := &c //copy of address
	c = 10

	fmt.Println(&c, d, &d) //address of c, address of c, address of d
	fmt.Println(c, *d)     // 10 10

	*d = 20
	fmt.Println(c, *d) // 20 20
}
