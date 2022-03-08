package main

import (
	"fmt"

	"github.com/go_practice/03_Packages_and_Imports/something"
)

func main() {
	fmt.Println("Hello world")

	something.SayHello() // Public function
	// something.sayBye() // Private function
}
