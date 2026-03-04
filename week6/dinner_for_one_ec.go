/***************************************
Dinner For One (extra credit version)
Author: Jeff Hewitt
Date Completed: 2 Mar 2026
Description: A program that functions as a food ordering service by
			 displaying menu options, taking your choice for meal
			 combo (including custom meal creation), offering a drink
			 and size upgrade, and displaying the final order summary at the end.
***************************************/

package main

import "fmt"

// custom type definition: "Combo"
type Combo struct {
	Entree string
	Side   string
	Drink  string
	Size   int
	Price  float64
}

// member function definition to display Combo option information
func (c Combo) Display() {
	fmt.Println("Entree:", c.Entree)
	fmt.Println("Side:", c.Side)
	fmt.Print("Price: $", c.Price, "\n")
}

// member function to alter Drink value for Combo objects
func (c *Combo) SetDrink(drinkInput string) {
	// keep default drink if not changed
	if drinkInput != "" {
		c.Drink = drinkInput
	}
}

// member function to alter Size value for Combo objects
func (c *Combo) SetSize(sizeInput int) {
	c.Size = sizeInput

	// alter order price if size is upgraded
	if sizeInput == 2 {
		c.Price += 2
	} else if sizeInput == 3 {
		c.Price += 3
	}
}

// member function to output the final order summary
func (c Combo) DisplayFull() {
	fmt.Println("Entree:", c.Entree)
	fmt.Println("Side:", c.Side)

	// translate integer size value to English (string)
	if c.Size == 1 {
		fmt.Println("Size: Small")
	} else if c.Size == 2 {
		fmt.Println("Size: Medium")
	} else {
		fmt.Println("Size: Large")
	}

	fmt.Println("Drink:", c.Drink)
	fmt.Print("Price: $", c.Price, "\n\n")
}

// pseudo-constructor for new "Combo" objects
func NewCombo(entreeInput string, sideInput string, drinkInput string, sizeInput int, priceInput float64) Combo {
	var c Combo
	c.Entree = entreeInput
	c.Side = sideInput
	c.Drink = drinkInput
	c.Size = sizeInput
	c.Price = priceInput
	return c
}

// non-member function to display menu and prompt user for combo choice
func PromptMenuAndSelectCombo() Combo {
	// instantiate variables to represent menu options
	var order1 Combo = NewCombo("Hamburger", "Fries", "Coke", 1, 5.99)
	var order2 Combo = NewCombo("Burrito", "Rice", "Coke", 1, 4.99)
	var order3 Combo = NewCombo("Salad", "Breadsticks", "Coke", 1, 4.49)

	var orderNum int
	var customEntree string
	var customSide string

	fmt.Println("Combo 1:")
	order1.Display()
	fmt.Println("\nCombo 2:")
	order2.Display()
	fmt.Println("\nCombo 3:")
	order3.Display()
	fmt.Println("\nCombo 4:")
	fmt.Println("Custom Order")
	fmt.Println("Price: $6.99")

	fmt.Println("\nPlease select your order number")
	fmt.Scanln(&orderNum)

	if orderNum == 1 {
		return order1
	} else if orderNum == 2 {
		return order2
	} else if orderNum == 3 {
		return order3
	} else {
		fmt.Println("Enter entree")
		fmt.Scanln(&customEntree)
		fmt.Println("Enter side")
		fmt.Scanln(&customSide)

		return NewCombo(customEntree, customSide, "Coke", 1, 6.99)
	}
}

// non-member function to prompt the user to upgrade their order size and drink choice
func UpdateSizeAndDrink(userCombo Combo) Combo {
	// instantiate variables to take Combo Size and Drink inputs
	var orderSize int
	var drinkChoice string

	fmt.Println("\nSize upgrade prices: Small $0, Medium $2, Large $3")
	fmt.Println("What size would you like? (1 = Small, 2 = Medium, 3 = Large)")
	fmt.Scanln(&orderSize)
	userCombo.SetSize(orderSize)

	fmt.Println("\nWhat would you like to drink? (leave blank to keep Coke as default)")
	fmt.Scanln(&drinkChoice)
	userCombo.SetDrink(drinkChoice)

	// return altered Combo object
	return userCombo
}

func main() {
	// declare variable for tracking the user's order (Combo object) throughout the scope of main()
	var userOrder Combo

	// display menu options and take the user's combo choice and assign choice to variable
	fmt.Print("Welcome to Eclectic Drive-Thru. What would you like to order?\n\n")
	userOrder = PromptMenuAndSelectCombo()
	//fmt.Println(userOrder) // debugging

	// prompt the user to upgrade their order size and drink choice
	userOrder = UpdateSizeAndDrink(userOrder)
	//fmt.Println(userOrder) // debugging

	// output order summary for the user
	fmt.Println("\nHere is your order:")
	userOrder.DisplayFull()
	fmt.Print("Thank you, please pull forward")
}
