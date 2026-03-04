package main

import "fmt"

func main() {
	// print/scan testing
	fmt.Println("Hola mundo")

	var foo int
	fmt.Println("Enter an integer value:")
	fmt.Scanln(&foo)
	fmt.Println(foo, "is half of", foo*2)

	// scope testing
	{
		var subScope int = 3
		fmt.Println(subScope)
	}

	fmt.Println(subScope)
}