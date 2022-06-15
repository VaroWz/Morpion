package main

import "fmt"

func checkWin() bool {

	//aretes check
	if placeLoc(1) == placeLoc(2) && placeLoc(1) == placeLoc(3) {
		if placeLoc(1) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(1) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(1) == placeLoc(4) && placeLoc(1) == placeLoc(7) {
		if placeLoc(1) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(1) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(7) == placeLoc(8) && placeLoc(7) == placeLoc(9) {
		if placeLoc(7) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(7) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(3) == placeLoc(6) && placeLoc(3) == placeLoc(9) {
		if placeLoc(3) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(3) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	//center check
	if placeLoc(2) == placeLoc(5) && placeLoc(5) == placeLoc(8) {
		if placeLoc(2) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(2) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(4) == placeLoc(5) && placeLoc(5) == placeLoc(6) {
		if placeLoc(4) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(4) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(1) == placeLoc(5) && placeLoc(5) == placeLoc(9) {
		if placeLoc(1) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(1) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(7) == placeLoc(5) && placeLoc(5) == placeLoc(3) {
		if placeLoc(7) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(7) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	return false
}
