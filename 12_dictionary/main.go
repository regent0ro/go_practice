package main

import (
	"fmt"
	"github/go_bank_dictionary/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	// definition, err := dictionary.Search("second") // errNotFound

	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)

	//err = dictionary.Add("first", "Greeting") // errWordExists
	err = dictionary.Add("Hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}

	definition, err = dictionary.Search("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("found Hello definition:", definition)

	// err = dictionary.Update("second", "Second word") // errCantUpdate
	err = dictionary.Update("first", "Second word")
	if err != nil {
		fmt.Println(err)
	}

	definition, err = dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("found first definition:", definition)

	// err = dictionary.Delete("hoge") // errNotFound
	err = dictionary.Delete("Hello")
	if err != nil {
		fmt.Println(err)
	}

	definition, err = dictionary.Search("Hello") // errnot found
	if err != nil {
		fmt.Println(err)
	}
}
