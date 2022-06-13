package checkWin

import "fmt"

var one = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var choose int
var isWin bool

func checkWin() {

	if one[0] == one[1] && one[0] == one[2] {
		if one[0] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[0] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		isWin = true
	}
	if one[0] == one[3] && one[0] == one[6] {
		if one[0] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[0] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		isWin = true
	}
	if one[6] == one[7] && one[6] == one[8] {
		if one[6] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[6] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		isWin = true
	}
	if one[2] == one[5] && one[2] == one[8] {
		if one[2] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[2] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		isWin = true
	}
}
