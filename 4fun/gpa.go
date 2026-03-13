package main

import "fmt"

//-------------------------------

type Class struct {
	Grade float64
	Units float64
}

func (c *Class) SetGrade(grade float64) {
	c.Grade = grade
}

func (c *Class) SetUnits(units float64) {
	c.Units = units
}

func (c *Class) CalculateGPA() float64 {
	return (c.Units * c.Grade) / c.Units
}

func NewClass(grade float64, units float64) Class {
	var c Class
	c.Grade = grade
	c.Units = units
	return c
}

//-------------------------------

// type Student struct {
// 	Class Class
// }

//-------------------------------

func TakeFloatInput() float64 {
	var tempFloat float64
	fmt.Scanln(&tempFloat)
	return tempFloat
}

func main() {
	var UserClass Class = NewClass(0, 0)

	fmt.Print("Welcome to the GPA calculator!\n\n")

	fmt.Println("Enter grade: (A = 4, B = 3, C = 2...)")
	UserClass.SetGrade(TakeFloatInput())
	fmt.Println("Enter units:")
	UserClass.SetUnits(TakeFloatInput())

	fmt.Println("GPA:", UserClass.CalculateGPA())
}
