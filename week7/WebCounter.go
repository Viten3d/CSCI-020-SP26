/****************************************************************
Web Counter
Author: Jeff Hewitt
Date Completed: 5 Mar 2026
Description: A program that tests a "hit counter's" functionality
			 for tracking, calling, and resetting the number visits
			 to a potential webpage.
****************************************************************/

package main

import "fmt"

// "WebCounter" custom type definition
type WebCounter struct {
	count int
}

// "WebCounter" member functions
// return the current count
func (wc *WebCounter) Get() int {
	return wc.count
}

// increment the count
func (wc *WebCounter) Hit() {
	wc.count += 1
}

// reset the count
func (wc *WebCounter) Reset() {
	wc.count = 0
}

// pseudo-constructor for "WebCounter" type
func NewWebCounter(countInput int) WebCounter {
	var wc WebCounter
	wc.count = countInput
	return wc
}

func main() {

	var wc WebCounter = NewWebCounter(42)

	fmt.Println("Web Counter Tester:")
	fmt.Println("\nStart:", wc.Get())

	wc.Hit()
	wc.Hit()

	fmt.Println("Two Hits:", wc.Get())

	wc.Hit()

	fmt.Println("Three Hits:", wc.Get())

	wc.Reset()

	fmt.Println("Reset:", wc.Get())

	wc.Hit()
	wc.Hit()
	wc.Hit()
	wc.Hit()

	fmt.Println("Four Hits:", wc.Get())
}