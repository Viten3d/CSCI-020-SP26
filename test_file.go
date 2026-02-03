package main

import "fmt"

func main() {
	fmt.Println("Hola mundo")

	var foo int
	fmt.Println("Enter an integer value:")
	fmt.Scanln(&foo)
	fmt.Println(foo, "is half of", foo*2)
}