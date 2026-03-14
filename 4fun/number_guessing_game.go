/*
Number Guessing Game
13 Mar 2026
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(int64(time.Now().Nanosecond()))
}

//------------------------------------------

type Game struct {
	InputNum int
	RandomNum int
}

func (g *Game) SetInputNum(num int) {
	g.InputNum = num
}

func (g Game) GetInputNum() int {
	return g.InputNum
}

func (g *Game) SetRandomNum() {
	g.RandomNum = rand.Intn(100) + 1 // 1 - 100
}

func (g Game) GetRandomNum() int {
	return g.RandomNum
}

func (g *Game) RunIntroSequence() {
	fmt.Println("Guess a number between 1 and 100:")
	g.SetInputNum(GetInt())
	g.SetRandomNum()
}

func (g Game) DisplayOutcome() {
	fmt.Print("\nThe number is ", g.GetRandomNum(), "\n\n")

	if g.InputNum == g.RandomNum {
		fmt.Println(g.GetInputNum(), "is correct! You won!")
	} else {
		fmt.Println(g.GetInputNum(), "is incorrect. You lost.")
	}
}

//------------------------------------------

func NewGame(input int, random int) Game {
	var g Game
	g.InputNum = input
	g.RandomNum = random
	return g
}

func GetInt() int {
	var num int
	fmt.Scanln(&num)
	return num
}

//------------------------------------------

func main() {
	// instantiate main() object
	var CurrentGame Game = NewGame(0, 0)

	// get number from user + generate random number
	CurrentGame.RunIntroSequence()
	
	// display random number + game outcome
	CurrentGame.DisplayOutcome()
}