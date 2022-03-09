package main

import "fmt"

func canIDrink(age int) bool {
	if koreanage := age + 2; koreanage < 20 { // only use for if block. variable expression
		return false
	}
	return true
}

func switchCanIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDrink(18))
	fmt.Println(canIDrink(16))

	fmt.Println(switchCanIDrink(18))
	fmt.Println(switchCanIDrink(16))
	fmt.Println(switchCanIDrink(52))
}
