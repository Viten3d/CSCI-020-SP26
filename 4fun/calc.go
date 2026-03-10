/*
Calculator
10 Mar 2026
*/

package main

import "fmt"

type Expression struct {
	FirstTerm float64
	SecondTerm float64
	Operator string
}

func (e *Expression) SetFirstTerm(ft float64) {
	e.FirstTerm = ft
}

func (e Expression) GetFirstTerm() float64 {
	return e.FirstTerm
}

func (e *Expression) SetSecondTerm(st float64) {
	e.SecondTerm = st
}

func (e Expression) GetSecondTerm() float64 {
	return e.SecondTerm
}

func (e *Expression) SetOperator(op string) {
	e.Operator = op
}

func (e Expression) GetOperator() string {
	return e.Operator
}

func (e Expression) EvaluateExpression() float64 {
	if e.Operator == "+" {
		return e.FirstTerm + e.SecondTerm
	} else if e.Operator == "-" {
		return e.FirstTerm - e.SecondTerm
	} else if e.Operator == "*" {
		return e.FirstTerm * e.SecondTerm
	} else { // expecting "/"
		return e.FirstTerm / e.SecondTerm
	}
}

func NewExpression(ft float64, st float64, op string) Expression {
	var e Expression
	e.FirstTerm = ft
	e.SecondTerm = st
	e.Operator = op
	return e
}

func PromptFirstTerm() float64 {
	var ft float64
	fmt.Println("Enter the first term:")
	fmt.Scanln(&ft)
	return ft
}

func PromptSecondTerm() float64 {
	var st float64
	fmt.Println("Enter the second term:")
	fmt.Scanln(&st)
	return st
}

func PromptOperator() string {
	var op string
	fmt.Println("Enter the operator:")
	fmt.Scanln(&op)
	return op
}

func TranslateOperator(op string) string {
	if op == "+" {
		return " plus "
	} else if op == "-" {
		return  " minus "
	} else if op == "*" {
		return " times "
	} else { // expecting "/"
		return " divided by "
	}
}

func main() {
	// create expression object
	var currentExp Expression = NewExpression(0, 0, "+")

	// welcome user
	fmt.Print("Welcome to the calculator program!\n\n")

	// get terms and operator from user
	currentExp.SetFirstTerm(PromptFirstTerm())
	fmt.Println()
	currentExp.SetSecondTerm(PromptSecondTerm())
	fmt.Println()
	currentExp.SetOperator(PromptOperator())
	fmt.Println()

	// evaluate and output the result of the expression
	fmt.Print(currentExp.GetFirstTerm(), TranslateOperator(currentExp.GetOperator()), currentExp.GetSecondTerm())
	fmt.Println(" equals", currentExp.EvaluateExpression())
}