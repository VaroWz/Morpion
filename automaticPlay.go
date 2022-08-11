package main

import (
	"math/rand"
)

func automaticPlay() {

	//horizontale haut
	if placeLoc(1) == "X" && placeLoc(2) == "X" && placeLoc(3) != "O" {
		one[2] = "O"

	} else if placeLoc(1) == "X" && placeLoc(3) == "X" && placeLoc(2) != "O" {
		one[1] = "O"

	} else if placeLoc(2) == "X" && placeLoc(3) == "X" && placeLoc(1) != "O" {
		one[0] = "O"

		//horizontale centre
	} else if placeLoc(4) == "X" && placeLoc(5) == "X" && placeLoc(6) != "O" {
		one[5] = "O"

	} else if placeLoc(5) == "X" && placeLoc(6) == "X" && placeLoc(4) != "O" {
		one[3] = "O"
		//horizontale bas
	} else if placeLoc(7) == "X" && placeLoc(8) == "X" && placeLoc(9) != "O" {
		one[8] = "O"

	} else if placeLoc(7) == "X" && placeLoc(9) == "X" && placeLoc(8) != "O" {
		one[7] = "O"

	} else if placeLoc(8) == "X" && placeLoc(9) == "X" && placeLoc(7) != "O" {
		one[6] = "O"

		//verticale gauche
	} else if placeLoc(1) == "X" && placeLoc(4) == "X" && placeLoc(7) != "O" {
		one[6] = "O"

	} else if placeLoc(1) == "X" && placeLoc(7) == "X" && placeLoc(4) != "O" {
		one[3] = "O"

	} else if placeLoc(4) == "X" && placeLoc(7) == "X" && placeLoc(1) != "O" {
		one[0] = "O"

		//verticale centre
	} else if placeLoc(2) == "X" && placeLoc(5) == "X" && placeLoc(8) != "O" {
		one[7] = "O"

	} else if placeLoc(2) == "X" && placeLoc(8) == "X" && placeLoc(5) != "O" {
		one[4] = "O"

	} else if placeLoc(5) == "X" && placeLoc(8) == "X" && placeLoc(2) != "O" {
		one[1] = "O"

		//verticale droit
	} else if placeLoc(3) == "X" && placeLoc(6) == "X" && placeLoc(9) != "O" {
		one[8] = "O"

	} else if placeLoc(3) == "X" && placeLoc(9) == "X" && placeLoc(6) != "O" {
		one[5] = "O"

	} else if placeLoc(6) == "X" && placeLoc(9) == "X" && placeLoc(3) != "O" {
		one[2] = "O"

		//diago /
	} else if placeLoc(3) == "X" && placeLoc(5) == "X" && placeLoc(7) != "O" {
		one[6] = "O"

	} else if placeLoc(3) == "X" && placeLoc(7) == "X" && placeLoc(5) != "O" {
		one[4] = "O"

	} else if placeLoc(5) == "X" && placeLoc(7) == "X" && placeLoc(3) != "O" {
		one[2] = "O"

		//diago \
	} else if placeLoc(1) == "X" && placeLoc(5) == "X" && placeLoc(9) != "O" {
		one[8] = "O"

	} else if placeLoc(1) == "X" && placeLoc(9) == "X" && placeLoc(5) != "O" {
		one[4] = "O"

	} else if placeLoc(5) == "X" && placeLoc(9) == "X" && placeLoc(1) != "O" {
		one[0] = "O"

	} else {
		a := rand.Intn(9)
		if placeLoc(a-1) != "O" && placeLoc(a-1) != "X" {
			one[a] = "O"
		}
	}
}
