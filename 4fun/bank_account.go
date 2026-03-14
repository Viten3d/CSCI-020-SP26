/*
Bank Simulator
13 Mar 2026
*/

package main

import "fmt"

//------------------------------------------

type Account struct	{
	Checking int
	Savings int
	Pin int
}

func (a *Account) SetChecking(chk int) {
	a.Checking = chk
}

func (a *Account) SetSavings(sav int) {
	a.Savings = sav
}

func (a *Account) SetPin(pin int) {
	a.Pin = pin
}

//------------------------------------------

func NewAccount(chk int, sav int, pin int) Account {
	var na Account
	na.Checking = chk
	na.Savings = sav
	na.Pin = pin
	return na
}

func GetInt() int {
	var num int
	fmt.Scanln(&num)
	return num
}

//------------------------------------------

func main() {
	var UserAccount Account = NewAccount(0, 0, 0000)


}
