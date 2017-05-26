package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var Matrix [][]*Element
var SaluteIniziale int

func main() {
	SaluteIniziale = 50
	height, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("height not valid")
	}
	width, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic("width not valid")
	}
	Matrix = make([][]*Element, height)
	for i := range Matrix { // inizializzazione matrice
		Matrix[i] = make([]*Element, width)
		for j := range Matrix[i] {
			chose := rand.Intn(2) //scelta rando cibo bug o vuoto (null)
			switch chose {
			case 0:
				Matrix[i][j] = new(Element) // insetto
				Matrix[i][j].IsFood = false
				Matrix[i][j].Age = 0
				Matrix[i][j].Health = SaluteIniziale
			case 1:
				Matrix[i][j] = nil //vuota
			case 2:
				Matrix[i][j] = new(Element) // cibo
				Matrix[i][j].IsFood = true
				Matrix[i][j].Health = 5
			}
		}
	}

	fmt.Println(Matrix)
}

func muovi(h int, w int) { // h verticale, w orizzontale
	elemento := Matrix[h][w]
	if elemento == nil && elemento.IsFood {
		return
	}

}
