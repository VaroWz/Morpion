package main

import (
	"fmt"
)

func main() {

	checkWin.isWin = false
	for checkWin.compteur := 0; checkWin.compteur < 100; checkWin.compteur++ {

		if checkWin.isWin == false {
			fmt.Println("Placez votre pion à la case: ")
			board()
			fmt.Scan(&choose.checkWin)
			checkWin.choose = checkWin.choose - 1
			checkWin.one[checkWin.choose] = "X"
			checkWin()
		}
		if checkWin.isWin == true {
			board()
			fmt.Println("Fin de la partie")
			return
		}

	}

}
func board() {

	fmt.Println(checkWin.one[0], "|", checkWin.one[1], "|", checkWin.one[2])
	fmt.Println(checkWin.one[3], "|", checkWin.one[4], "|", checkWin.one[5])
	fmt.Println(checkWin.one[6], "|", checkWin.one[7], "|", checkWin.one[8])

}
