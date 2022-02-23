/*
  get N and output a sentence
  N입력을 받아서 문장 출력하기
  Nの入力をもらって、文章を出力する
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

	for i := 0; i <= n; i++ {
		fmt.Println("An apple is", i*10, "dollars, a book is", i*20, "dollars")
	}
}
