/***************************************
JOUST
Author: Jeff Hewitt
Date Started: 12 Apr 2026
Last Modified: 13 Apr 2026
Description: Simulate a set amount of jousting matches between
			 knights with with parameters set by command line
			 arguments, then display the win rate for each knight.
***************************************/

///----------
/// Preamble:

package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// random seed definition
func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

///------------------------
/// Knight Type Definition:

type Knight struct {
	Name    string
	Stamina int
	Weapon  Weapon
	Mounted bool
	Wins    int
	Draws   int
	Matches int
}

func (k *Knight) SetStamina(stam int) {
	k.Stamina = stam
}

func (k *Knight) HitWins() {
	k.Wins++
}

func (k *Knight) HitDraws() {
	k.Draws++
}

func (k *Knight) HitMatches() {
	k.Matches++
}

func (k *Knight) Joust() bool {
	k.Stamina -= k.Weapon.GetStaminaCost()
	return k.Weapon.Swing()
}

func (k *Knight) SetMounted(hitState bool) {
	k.Mounted = hitState
}

func (k Knight) GetMounted() bool {
	return k.Mounted
}

func (k Knight) GetStamina() int {
	return k.Stamina
}

func (k Knight) GetName() string {
	return k.Name
}

func (k Knight) GetWeapon() Weapon {
	return k.Weapon
}

func (k Knight) DisplayStats() {
	fmt.Println(
		k.Name, "with", k.Weapon.GetType(),
		"{ S:", k.Stamina,
		"M:", k.Mounted, "}",
	)
}

func (k Knight) DisplayMatchData() {
	fmt.Print(
		k.Name, " with ", k.Weapon.GetType(),
		"\nInitial Stamina: ", k.Stamina,
		"\nStamina Cost: ", k.Weapon.GetStaminaCost(),
		"\nHit Chance: ", k.Weapon.GetHitChance(),
		"\nWins: ", k.Wins,
		"\nDraws: ", k.Draws,
		//"\nMatches: ", k.Matches,
		"\nWin Rate: ", CalculateRate(k.Wins, k.Matches), "%",
		"\nDraw Rate: ", CalculateRate(k.Draws, k.Matches), "%\n",
	)
}

// pseudo-constructor
func NewKnight(name string, stam int, weap Weapon) Knight {
	var k Knight
	k.Name = name
	k.Stamina = stam
	k.Weapon = weap
	k.Mounted = true // default
	k.Wins = 0       // default
	k.Draws = 0      // default
	k.Matches = 0    // default
	return k
}

///------------------------
/// Weapon Type Definition:

type Weapon struct {
	StaminaCost int
	HitChance   int
	Type        string
}

func (w Weapon) GetStaminaCost() int {
	return w.StaminaCost
}

func (w Weapon) GetType() string {
	return w.Type
}

func (w Weapon) GetHitChance() int {
	return w.HitChance
}

func (w Weapon) Swing() bool {
	var randNum int = rand.Intn(100) + 1

	if randNum <= w.HitChance {
		return true
	} else {
		return false
	}
}

// pseudo-constructor
func NewWeapon(stamCos int, hitChn int, typ string) Weapon {
	var w Weapon
	w.StaminaCost = stamCos
	w.HitChance = hitChn
	w.Type = typ
	return w
}

///----------------------
/// Non-Member Functions:

func PostStep(k1M bool, k1S int, k2M bool, k2S int) bool {
	if k1M == false || k2M == false {
		return true
	} else if k1S <= 0 || k2S <= 0 {
		return true
	} else {
		return false
	}
}

func DetermineOutcome(k1M bool, k1S int, k1N string, k2M bool, k2S int, k2N string) int {
	var gState int // 0 = tie; 1 = K1 win; 2 = K2 win

	if k1M == true && k1S > 0 { // k1 win
		if k2M == false || k2S <= 0 {
			gState = 1
		}
	} else if k2M == true && k2S > 0 { // k2 win
		if k1M == false || k1S <= 0 {
			gState = 2
		}
	} else { // draw
		gState = 0
	}

	return gState
}

func ShowOutcomeText(k1N string, k2N string, gs int) {
	if gs == 1 { // k1 win
		fmt.Println(k1N, "wins!")
	} else if gs == 2 { // k2 win
		fmt.Println(k2N, "wins!")
	} else { // draw
		fmt.Println("It's a draw!")
	}
}

func ShowRoundData(k1 Knight, k2 Knight) {
	k1.DisplayStats()
	k2.DisplayStats()
	fmt.Println()
}

func CalculateRate(num int, den int) float64 {
	var rate float64
	rate = float64(num) / float64(den)
	rate *= 10000
	rate = math.Round(rate)
	rate /= 100
	return rate
}

func main() {
	// other variable declaration/instantiation
	var endGame bool = false
	var gameState int             // 0 = tie; 1 = K1 win; 2 = K2 win
	var matchCountDefault int = 5 // for convenience

	// declare variables for command line arguments
	var showRoundLogs bool
	var showMatchLogs bool
	var matchCount int

	var Knight1Name string
	var Knight1Stamina int
	var Knight1WeaponType string
	var Knight1HitChance int
	var Knight1StaminaCost int

	var Knight2Name string
	var Knight2Stamina int
	var Knight2WeaponType string
	var Knight2HitChance int
	var Knight2StaminaCost int

	// acquire and assign command line arguments
	flag.BoolVar(&showRoundLogs, "ShowRoundLogs", false, "Display knight and weapon statistics from each round as the program runs. (default false)")
	flag.BoolVar(&showMatchLogs, "ShowMatchLogs", false, "Display match winner as the program runs. (default false)")
	flag.IntVar(&matchCount, "Matches", matchCountDefault, "Amount of matches to be simulated.")

	flag.StringVar(&Knight1Name, "Knight1Name", "King Arthur", "Name of first knight.")
	flag.IntVar(&Knight1Stamina, "Knight1Stamina", 50, "Max stamina for first knight.")
	flag.StringVar(&Knight1WeaponType, "Knight1WeaponType", "Excalibur", "First knight's weapon name.")
	flag.IntVar(&Knight1HitChance, "Knight1HitChance", 15, "Hit chance of first knight's weapon.")
	flag.IntVar(&Knight1StaminaCost, "Knight1StaminaCost", 10, "Stamina cost of first knight's weapon.")

	flag.StringVar(&Knight2Name, "Knight2Name", "Black Knight", "Name of second knight.")
	flag.IntVar(&Knight2Stamina, "Knight2Stamina", 40, "Max stamina for second knight.")
	flag.StringVar(&Knight2WeaponType, "Knight2WeaponType", "Longsword", "Second knight's weapon name.")
	flag.IntVar(&Knight2HitChance, "Knight2HitChance", 10, "Hit chance of second knight's weapon.")
	flag.IntVar(&Knight2StaminaCost, "Knight2StaminaCost", 5, "Stamina cost of second knight's weapon.")

	flag.Parse()

	// create Knight and Weapon objects
	var Knight1 Knight = NewKnight(
		Knight1Name,
		Knight1Stamina,
		NewWeapon(Knight1StaminaCost, Knight1HitChance, Knight1WeaponType),
	)

	var Knight2 Knight = NewKnight(
		Knight2Name,
		Knight2Stamina,
		NewWeapon(Knight2StaminaCost, Knight2HitChance, Knight2WeaponType),
	)

	for iter := 0; iter < matchCount; iter++ {
		if showRoundLogs == true {
			fmt.Print("Round ", iter+1, ":\n")
		}

		for endGame == false {
			// run jousting simulation
			Knight2.SetMounted(!Knight1.Joust())
			Knight1.SetMounted(!Knight2.Joust())

			// display stats
			if showRoundLogs == true {
				ShowRoundData(Knight1, Knight2)
			}

			// check if game-ending conditions are met
			endGame = PostStep(
				Knight1.GetMounted(),
				Knight1.GetStamina(),
				Knight2.GetMounted(),
				Knight2.GetStamina(),
			)
		}

		// update wins, draws, and matches for knights
		Knight1.HitMatches()
		Knight2.HitMatches()

		gameState = DetermineOutcome(
			Knight1.GetMounted(),
			Knight1.GetStamina(),
			Knight1.GetName(),
			Knight2.GetMounted(),
			Knight2.GetStamina(),
			Knight2.GetName(),
		)

		if gameState == 1 {
			Knight1.HitWins()
		} else if gameState == 2 {
			Knight2.HitWins()
		} else {
			Knight1.HitDraws()
			Knight2.HitDraws()
		}

		if showMatchLogs == true {
			ShowOutcomeText(Knight1.GetName(), Knight2.GetName(), gameState)
		}

		if showRoundLogs == true || showMatchLogs == true {
			fmt.Println()
		}

		// reset values for the next match
		//if iter != matchCount-1 {
		endGame = false
		Knight1.SetStamina(Knight1Stamina)
		Knight2.SetStamina(Knight2Stamina)
		//}
	}

	fmt.Print("Total Matches: ", matchCount, "\n\n")
	fmt.Println("Knight 1:")
	Knight1.DisplayMatchData()
	fmt.Println()
	fmt.Println("Knight 2:")
	Knight2.DisplayMatchData()
}
