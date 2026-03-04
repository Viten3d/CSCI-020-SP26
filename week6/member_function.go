/***************************************
Member Function
Author: Jeff Hewitt
Date Completed: 28 Feb 2026
Description: Outputs the member variable values of a preconstructed custom object using a member function.
***************************************/

package main

import "fmt"

// custom type definition for "Combo"
type Combo struct {
	Entree string
	Side   string
	Price  float64
}

// pseudo-constructor for "Combo" objects
func NewCombo(ent string, sid string, pri float64) Combo {
	var com Combo
	com.Entree = ent
	com.Side = sid
	com.Price = pri
	return com
}

// member function definition for "Display" (value-receiver)
func (c Combo) Display() {
	fmt.Println("Entree:", c.Entree)
	fmt.Println("Side:", c.Side)
	fmt.Print("Price: $", c.Price, "\n")
}

func main() {
	// instantiate "combo" objects using pseudo-constructor
	var order1 Combo = NewCombo("Hamburger", "Fries", 5.99)
	var order2 Combo = NewCombo("Burrito", "Rice", 4.99)
	var order3 Combo = NewCombo("Salad", "Breadsticks", 4.49)

	// output object elements using "Display" member function
	fmt.Println("Combo 1:")
	order1.Display()
	fmt.Println("\nCombo 2:")
	order2.Display()
	fmt.Println("\nCombo 3:")
	order3.Display()
}
