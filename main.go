package main

import (
	"fmt"
)

var one = []string{"?", "?", "?", "?", "?", "?", "?", "?", "?"}
var choose int
var isWin bool

func main() {

	isWin = false
	for compteur := 0; compteur < 100; compteur++ {

		if isWin == false {
			fmt.Println("Placez votre pion à la case: ")
			board()
			fmt.Scan(&choose)
			choose = choose - 1
			one[choose] = "X"
			checkWin()
		}
		if isWin == true {
			board()
			fmt.Println("Fin de la partie")
			return
		}

	}

}
func board() {

	fmt.Println(one[0], "|", one[1], "|", one[2])
	fmt.Println(one[3], "|", one[4], "|", one[5])
	fmt.Println(one[6], "|", one[7], "|", one[8])

}

func checkWin() {

	if one[0] == one[1] && one[0] == one[2] || one[0] == one[3] && one[0] == one[6] || one[6] == one[7] && one[6] == one[8] || one[2] == one[5] && one[2] == one[8] {
		if one[0] == "X" {
			fmt.Println("Vous avez gagné !!!")
		}
		if one[0] == "O" {
			fmt.Println("Vous avez perdu !")
		}
		isWin = true
	}
}
