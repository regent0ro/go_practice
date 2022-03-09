package main

import "fmt"

func main() {
	names := [5]string{"hoge", "hige", "hige"}

	names[3] = "huge"
	names[4] = "hege"
	//names[5] = "hhge" //error
	fmt.Println(names)

	names2 := []string{"hoge", "hige", "hige"}

	// names2[3] = "huge" //error
	names2 = append(names2, "huge", "hege") // return append array
	fmt.Println(names2)
}
