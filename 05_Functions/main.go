package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//using variable arguments
func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return, defer
func nakedLenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("nakedLenAndUpper is done")
	fmt.Println("nakedLenAndUpper is started")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	fmt.Println(multiply(2, 2)) // print 4

	totalLength, upperName := lenAndUpper("hoge")
	fmt.Println(totalLength, upperName) // print 4, HOGE

	totalLength2, _ := lenAndUpper("hige") // ignore upperName value
	fmt.Println(totalLength2)              // print 4

	repeatMe("hoge", "hage", "hige", "huge", "hege") //using variable arguments

	nakedTotalLength, nakedUpperName := nakedLenAndUpper("hoge")
	fmt.Println(nakedTotalLength, nakedUpperName) // print 4, HOGE

}
