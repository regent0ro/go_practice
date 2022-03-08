package main

import "fmt"

func main() {
	const name string = "hoge"
	//name = "hige" // const error
	fmt.Println(name)

	var name2 string = "hige"
	name2 = "huge"
	fmt.Println(name2)

	name3 := "hage" // go recognize type
	name3 = "hege"
	fmt.Println(name3)
}
