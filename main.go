package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//VARIABILI
var Matrix [][]*Element
var Altezza int
var Larghezza int
var SaluteIniziale int
var Clock uint
var NumClock uint

func main() { //FUNZIONE MAIN
	SaluteIniziale = 50
	Clock = 1
	NumClock = 0
	rand.Seed(time.Now().UTC().UnixNano()) //inizializzazione rand
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
			chose := rand.Intn(3) //scelta rando cibo bug o vuoto (null)
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

	aggiorna()
}

func aggiorna() { //FUNZIONE AGGIORNA:	chiama la funzione muovi
	for {
		time.Sleep(time.Second * time.Duration(Clock))
		NumClock++
		for i := 0; i < Altezza; i++ {
			for j := 0; j < Larghezza; j++ {
				muovi(i, j)
			}
		}
		fmt.Printf("\nSituazione dopo %d movimenti:\n", NumClock)
		stampaMatrice()
	}
}

func muovi(h int, w int) { //FUNZIONE MUOVI:	aggiorna la posizione di tutti gli oggetti in tabella	// h verticale, w orizzontale
	elemento := Matrix[h][w]                //assegnamente del contenuto della cella in 'elemento'
	if elemento == nil || elemento.IsFood { //controllo se 'elemento' è cibo o un altro essere
		return
	}
	direzCasOriz := rand.Intn(3)	//numero da 0 a 2
	direzCasOriz--
	direzCasVert := rand.Intn(3)
	direzCasVert--
	nuovaPosizioneH := h + direzCasVert //aggiornamento posiozione verticale
	nuovaPosizioneW := w + direzCasOriz //aggiornamento posizione orizzontale

	if nuovaPosizioneH >= Altezza || nuovaPosizioneH < 0 ||
		nuovaPosizioneW >= Larghezza || nuovaPosizioneW < 0 { //se esce dai bordi
		return
	}

	if tmpNewElem := Matrix[nuovaPosizioneH][nuovaPosizioneW]; tmpNewElem != nil {

		if tmpNewElem.Razza!=elemento.Razza {		//se non è dalla stessa razza
			if tmpNewElem.IsFood || (tmpNewElem.Health+tmpNewElem.Evoluzione) < (elemento.Health+elemento.Evoluzione) { // se e' cibo o un insetto piu debole
				elemento.Health += tmpNewElem.Health	//prelevamento energia essere fagocitato
				Matrix[nuovaPosizioneH][nuovaPosizioneW] = elemento	//inglobamento essere peritos
			} else {
				Matrix[h][w] = nil	//perdita nel combattimento per la sopravvivenza
				tmpNewElem.Health += elemento.Health	//il nemico prende l'energia
			}
		} else {	//se sono amici
			if nuovaPosizioneH==h && nuovaPosizioneW==w {	//se cerca di mangiare il suo amico
				muovi(h, w)
			}
		}
	} else { //si muove sulla nuova casella
		Matrix[nuovaPosizioneH][nuovaPosizioneW] = elemento
		Matrix[h][w] = nil
		elemento.Health-=elemento.CostoMov

		if rand.Intn(10)==0 {	//se ha fortuna (o sfortuna) si evolve
			if rand.Intn(3)==0 {
				elemento.Evoluzione--
			} else {
				elemento.Evoluzione++
			}
		}

		if (elemento.Health-elemento.Premura)>elemento.CostoSex {		//se ha energia a sufficienza per riprodursi
			Matrix[h][w] = Costruttore(elemento.Razza, elemento.Evoluzione, elemento.CostoMov, elemento.CostoSex, elemento.Premura)
		}

	}
}

func stampaMatrice2() {
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

func stampaMatrice() {
	for i := 0; i < Altezza; i++ {
		for j := 0; j < Larghezza; j++ {
			if Matrix[i][j]!=nil {
				fmt.Printf("0 ")
			} else {
				fmt.Printf("1 ")
			}
			fmt.Printf("  Colonna %d: %s\n")
		}
		fmt.Printf("\n")
	}
}
