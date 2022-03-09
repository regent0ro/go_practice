package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	//maps
	user := map[string]string{"name": "hoge", "age": "12"}
	fmt.Println(user)

	for key, value := range user {
		fmt.Println(key, value)
	}

	//structs
	person_one := person{"hoge", 12, []string{"pizza", "pasta"}}
	fmt.Println(person_one)
	fmt.Println(person_one.favFood)

	person_two := person{name: "hoge", age: 12, favFood: []string{"pizza", "pasta"}}
	fmt.Println(person_two)
}
