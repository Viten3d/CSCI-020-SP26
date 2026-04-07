/***************************************
HiLo
Author: Jeff Hewitt
Date Created: 09 Feb 2026
Date Modified: 21 Feb 2026
Description: A program that lets you guess whether a number will be greater or less than the previous.
TBD:
- add bet system
- add continuous or preset round quantity
- prevent incorrect inputs
	- bet amount
	- guess choice
- prevent next number duplicates
***************************************/

// Preamble.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// define random seed argument
func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

// Player type definition + pseudo-constructor.
type Player struct {
	Guess       int
	CoinBalance int
	Bet         int
	Score       int
}

func (p *Player) SetBet(bet int) {
	p.Bet = bet
}

func (p *Player) SetGuess(g int) {
	p.Guess = g
}

func (p *Player) WithdrawBet() {
	p.CoinBalance -= p.Bet
}

func (p *Player) IncrementScore() {
	p.Score += 1
}

func (p *Player) WinBet() {
	p.CoinBalance += p.Bet * 2
}

func (p Player) DisplayStats() {
	fmt.Println("Current balance:", p.CoinBalance)
	fmt.Println("Current score:", p.Score)
}

func NewPlayer(g int, bet int) Player {
	var p Player
	p.Guess = g
	p.CoinBalance = 1000 // default
	p.Bet = bet
	p.Score = 0 // default
	return p
}

// Game type definition + pseudo-constructor.
type Game struct {
	RangeUpperLimit int // 0 - n
	CurrentNumber   int
	NextNumber      int
}

func (g *Game) SetNumbers() {
	for g.CurrentNumber == g.NextNumber { // ensure two different numbers
		g.CurrentNumber = rand.Intn(g.RangeUpperLimit + 1)
		g.NextNumber = rand.Intn(g.RangeUpperLimit + 1)
	}
}

func (g Game) GetCurrentNumber() int {
	return g.CurrentNumber
}

func (g Game) GetNextNumber() int {
	return g.NextNumber
}

func NewGame(upperlim int) Game {
	var g Game
	g.RangeUpperLimit = upperlim
	g.CurrentNumber = 0 // default
	g.NextNumber = 0    // default
	return g
}

// Non-member functions
func PromptRangeUpperLimit() int {
	fmt.Println("Enter the upper limit for the number range:")
	return TakeInt()
}

func DisplayCurrentNumber(g Game) {
	fmt.Println("Current Number:", g.GetCurrentNumber())
}

func PromptGuess() int {
	fmt.Println("Enter your guess (0 = lower; 1 = higher):")
	return TakeInt()
}

func PromptBet() int {
	fmt.Println("Enter the amount you'd like to bet:")
	return TakeInt()
}

func TakeInt() int {
	var n int
	fmt.Scanln(&n)
	return n
}

func DisplayGameResult(gstate int) {
	if gstate == 1 {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost.")
	}
}

func DetermineGameState(g Game, p Player) int {
	var gameState int
	if p.Guess == 0 && g.NextNumber < g.CurrentNumber {
		gameState = 1
	} else if p.Guess == 1 && g.NextNumber > g.CurrentNumber {
		gameState = 1
	} else {
		gameState = 0
	}
	return gameState
}

func main() {

	// welcome message
	fmt.Println("Welcome to the Hi-Lo Game!")
	fmt.Println("")

	// instantiate Game and Player object + variable to track the game result
	var GameState int
	var CurrentPlayer Player = NewPlayer(0, 0)
	var CurrentGame Game = NewGame(PromptRangeUpperLimit())
	fmt.Println()

	// generate numbers
	CurrentGame.SetNumbers()

	// show current number
	DisplayCurrentNumber(CurrentGame)

	// update user's guess and bet
	CurrentPlayer.SetGuess(PromptGuess())
	CurrentPlayer.SetBet(PromptBet())

	// process bet transaction + show next number
	CurrentPlayer.WithdrawBet()
	fmt.Println("The next number is:", CurrentGame.GetNextNumber())

	// calculate game results and apply necessary rewards
	GameState = DetermineGameState(CurrentGame, CurrentPlayer)

	if GameState == 1 {
		CurrentPlayer.IncrementScore()
		CurrentPlayer.WinBet()
	}

	// show game results and user stats
	DisplayGameResult(GameState)
	CurrentPlayer.DisplayStats()
}
