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
var SaluteIniziale int = 50
var CostoMovIniziale int = 5
var CostoSexIniziale int = 100
var EvoluzioneIniziale int = 0
var PremuraIniziale int = 10
var AgeLimite int = 30
var Clock uint
var NumClock uint

func main() { //FUNZIONE MAIN
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
				Matrix[i][j].CostoMov = CostoMovIniziale
				Matrix[i][j].CostoSex = CostoSexIniziale
				Matrix[i][j].Evoluzione = EvoluzioneIniziale
				Matrix[i][j].Premura = PremuraIniziale
				Matrix[i][j].Razza = rand.Intn(2)
			case 1:
				Matrix[i][j] = nil //vuota
			case 2:
				Matrix[i][j] = new(Element) // cibo
				Matrix[i][j].IsFood = true
				Matrix[i][j].Health = 5
			}
		}
	}

	go ServiHTML() // fai partire il server html

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

	if elemento.Health<=0 {
		Matrix[h][w] = nil
		return
	}

	if elemento.Age>AgeLimite {
		Matrix[h][w] = nil
		return
	}
	direzCasOriz := rand.Intn(3) //numero da 0 a 2
	direzCasOriz--
	direzCasVert := rand.Intn(3)
	direzCasVert--
	nuovaPosizioneH := h + direzCasVert //aggiornamento posiozione verticale
	nuovaPosizioneW := w + direzCasOriz //aggiornamento posizione orizzontale

	if nuovaPosizioneH >= Altezza || nuovaPosizioneH < 0 ||
		nuovaPosizioneW >= Larghezza || nuovaPosizioneW < 0 { //se esce dai bordi
		return
	}

	if Matrix[nuovaPosizioneH][nuovaPosizioneW] != nil {
		if Matrix[nuovaPosizioneH][nuovaPosizioneW].Razza != Matrix[h][w].Razza { //se non è dalla stessa razza
			if Matrix[nuovaPosizioneH][nuovaPosizioneW].IsFood || (Matrix[nuovaPosizioneH][nuovaPosizioneW].Health+Matrix[nuovaPosizioneH][nuovaPosizioneW].Evoluzione) < (Matrix[h][w].Health+elemento.Evoluzione) { // se e' cibo o un insetto piu debole
				Matrix[h][w].Health += Matrix[nuovaPosizioneH][nuovaPosizioneW].Health                //prelevamento energia essere fagocitato
				Matrix[nuovaPosizioneH][nuovaPosizioneW] = Matrix[h][w] //inglobamento essere perito
				Matrix[h][w] = nil
				Matrix[nuovaPosizioneH][nuovaPosizioneW].Health -= Matrix[nuovaPosizioneH][nuovaPosizioneW].CostoMov
				Matrix[nuovaPosizioneH][nuovaPosizioneW].Age++
			} else {	//perdita nel combattimento per la sopravvivenza
				Matrix[nuovaPosizioneH][nuovaPosizioneW].Health += Matrix[h][w].Health //il nemico prende l'energia
				Matrix[h][w] = nil
			}
		} else { //se sono amici
			if nuovaPosizioneH == h && nuovaPosizioneW == w { //se cerca di mangiare il suo amico
				muovi(h, w)
			}
		}
	} else { //si muove sulla nuova casella
		Matrix[nuovaPosizioneH][nuovaPosizioneW] = Matrix[h][w]
		Matrix[nuovaPosizioneH][nuovaPosizioneW].Health -= Matrix[nuovaPosizioneH][nuovaPosizioneW].CostoMov
		Matrix[nuovaPosizioneH][nuovaPosizioneW].Age++
		Matrix[h][w] = nil

		if rand.Intn(10) == 0 { //se ha fortuna (o sfortuna) si evolve
			if rand.Intn(3) == 0 {
				Matrix[nuovaPosizioneH][nuovaPosizioneW].Evoluzione--
			} else {
				Matrix[nuovaPosizioneH][nuovaPosizioneW].Evoluzione++
			}
		}

		if (Matrix[nuovaPosizioneH][nuovaPosizioneW].Health-Matrix[nuovaPosizioneH][nuovaPosizioneW].Premura)>Matrix[nuovaPosizioneH][nuovaPosizioneW].CostoSex {		//se ha energia a sufficienza per riprodursi
			Matrix[h][w] = Costruttore(Matrix[nuovaPosizioneH][nuovaPosizioneW].Razza, Matrix[nuovaPosizioneH][nuovaPosizioneW].Evoluzione, Matrix[nuovaPosizioneH][nuovaPosizioneW].CostoMov, Matrix[nuovaPosizioneH][nuovaPosizioneW].CostoSex, Matrix[nuovaPosizioneH][nuovaPosizioneW].Premura, SaluteIniziale)
		}

	}
}

func stampaMatrice() {
	for i := 0; i < Altezza; i++ {
		for j := 0; j < Larghezza; j++ {
			if Matrix[i][j] == nil {
				fmt.Printf("    --  ")
			} else {
				if Matrix[i][j].IsFood {
					fmt.Printf("    CC  ")
				} else {
					fmt.Printf("%d   %d  ",Matrix[i][j].Razza, Matrix[i][j].Health)
				}
			}
		}
		fmt.Printf("\n")
		
	}
}
