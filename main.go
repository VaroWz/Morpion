package main

import (
	"fmt"
)

var one = []string{"?", "?", "?", "?", "?", "?", "?", "?", "?"}
var choose int

func main() {

	for compteur := 0; compteur < 100; compteur++ {

		fmt.Println("Placez votre pion à la case: ")
		board()
		fmt.Scan(&choose)
		choose = choose - 1
		one[choose] = "X"
		checkWin()
	}

}
func board() {

	fmt.Println(one[0], "|", one[1], "|", one[2])
	fmt.Println(one[3], "|", one[4], "|", one[5])
	fmt.Println(one[6], "|", one[7], "|", one[8])

}

func checkWin() {

	if one[0] == one[1] && one[0] == one[3] {
		fmt.Println("Le jeu est terminé")
	}
}
