package main

import "fmt"

func checkWin() bool {

	//aretes check
	if placeLoc(0) == placeLoc(1) && placeLoc(0) == placeLoc(2) {
		if placeLoc(0) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(0) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(0) == placeLoc(3) && placeLoc(0) == placeLoc(6) {
		if placeLoc(0) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(0) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(6) == placeLoc(7) && placeLoc(6) == placeLoc(8) {
		if placeLoc(6) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(6) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(2) == placeLoc(5) && placeLoc(2) == placeLoc(8) {
		if placeLoc(2) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(2) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	//center check
	if placeLoc(1) == placeLoc(4) && placeLoc(4) == placeLoc(7) {
		if placeLoc(1) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(1) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(3) == placeLoc(4) && placeLoc(4) == placeLoc(5) {
		if placeLoc(3) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(3) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(0) == placeLoc(4) && placeLoc(4) == placeLoc(8) {
		if placeLoc(0) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(0) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if placeLoc(6) == placeLoc(4) && placeLoc(4) == placeLoc(2) {
		if placeLoc(6) == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if placeLoc(6) == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	return false
}
