package main

import (
	"fmt"
)

var one = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var choose int

func main() {

	for compteur := 0; compteur < 100; compteur++ {

		fmt.Println("Placez votre pion à la case: ")
		board()
		fmt.Scan(&choose)
		choose = choose - 1
		if one[choose] == "X" || one[choose] == "O"{
			fmt.Println("Vous ne pouvez pas jouer ici !")
		} else{

			one[choose] = "X"
			if checkWin() == true {
				board()
				fmt.Println("Fin de la partie")
				return
			} else {
				automaticPlay()
			}
		}
	}

}
func board() {

	fmt.Println(one[0], "|", one[1], "|", one[2])
	fmt.Println(one[3], "|", one[4], "|", one[5])
	fmt.Println(one[6], "|", one[7], "|", one[8])

}
