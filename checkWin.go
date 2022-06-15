package main

import "fmt"

func checkWin() bool {

	//aretes check
	if one[0] == one[1] && one[0] == one[2] {
		if one[0] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[0] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if one[0] == one[3] && one[0] == one[6] {
		if one[0] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[0] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if one[6] == one[7] && one[6] == one[8] {
		if one[6] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[6] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if one[2] == one[5] && one[2] == one[8] {
		if one[2] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[2] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	//center check
	if one[1] == one[4] && one[4] == one[7] {
		if one[1] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[1] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if one[3] == one[4] && one[4] == one[5] {
		if one[3] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[3] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if one[0] == one[4] && one[4] == one[8] {
		if one[0] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[0] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	if one[6] == one[4] && one[4] == one[2] {
		if one[6] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[6] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		return true
	}
	return false
}
