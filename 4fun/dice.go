/*
Dice Roller
30 Mar 2026
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
	"flag"
)

func init() {
    rand.Seed(int64(time.Now().Nanosecond()))
}

type Dice struct {
	Sides int
	RollValue int
	Qty int
}

func (d Dice) GetSides() int {
	return d.Sides
}

func (d Dice) GetRollValue() int {
	return d.RollValue
}

func (d Dice) GetQty() int {
	return d.Qty
}

// setter for Dice.RollValue
func (d *Dice) Roll() {
	d.RollValue = rand.Intn(d.Sides) + 1
}

// pseudo-constructor for Dice objects
func NewDice(s int, rv int, q int) Dice {
	var d Dice
	d.Sides = s
	d.RollValue = rv
	d.Qty = q
	return d
}

func main() {
	var userSides, userQty int

	flag.IntVar(&userSides, "Sides", 6, "Enter the number of sides.")
	flag.IntVar(&userQty, "Qty", 1, "Enter the quantity of dice.")
	flag.Parse()

	var userDice Dice = NewDice(userSides, 0, userQty)

	fmt.Print(userDice.GetQty(), " D", userDice.GetSides())
	fmt.Println(" roll(s):")

	for i := 0; i < userDice.Qty; i += 1 {
		userDice.Roll()
		fmt.Println(userDice.GetRollValue())
	}
}