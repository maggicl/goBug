package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var Matrix [][]*Element
var Altezza int
var Larghezza int
var SaluteIniziale int
var Clock uint

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
	Altezza = height
	Larghezza = width
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
	fmt.Println("Situazione iniziale: ")
	stampaMatrice()

	//go aggiorna()
}

func aggiorna() {
	for {

		time.Sleep(time.Second * time.Duration(Clock))
	}
}

func muovi(h int, w int) { // h verticale, w orizzontale
	elemento := Matrix[h][w]
	if elemento == nil || elemento.IsFood {
		return
	}
	direzCasOriz := rand.Intn(2)
	direzCasOriz--
	direzCasVert := rand.Intn(2)
	direzCasVert--
	nuovaPosizioneH := h + direzCasVert
	nuovaPosizioneW := w + direzCasOriz
	if nuovaPosizioneH > Altezza || nuovaPosizioneH < 0 {
		muovi(h, w)
	}

	if nuovaPosizioneW > Larghezza || nuovaPosizioneW < 0 {
		muovi(h, w)
	}

	if tmpNewElem := Matrix[nuovaPosizioneH][nuovaPosizioneW]; tmpNewElem != nil {
		if tmpNewElem.IsFood || tmpNewElem.Health < elemento.Health { // se e' cibo o un insetto piu debole
			elemento.Health += tmpNewElem.Health
			Matrix[nuovaPosizioneH][nuovaPosizioneW] = elemento
		} else {
			Matrix[h][w] = nil
			tmpNewElem.Health += elemento.Health
		}
	} else {
		Matrix[nuovaPosizioneH][nuovaPosizioneW] = elemento
		Matrix[h][w] = nil
	}
}

func stampaMatrice() {
	for i := 0; i < Altezza; i++ {
		fmt.Printf("Riga %d:\n", i)
		for j := 0; j < Larghezza; j++ {
			var stringa string
			elem := Matrix[i][j]
			if elem == nil {
				stringa = "Vuota"
			} else {
				stringa = elem.String()
			}
			fmt.Printf("  Colonna %d: %s\n", j, stringa)
		}
	}
}
