package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"os/exec"
)

//VARIABILI
var def int
var Matrix [][]*Element
var Altezza int
var Larghezza int
var SaluteIniziale [2]int
var CostoMovIniziale [2]int
var CostoSexIniziale [2]int
var EvoluzioneIniziale [2]int
var PremuraIniziale int = 10
var AgeMaxInizio [2]int
var Clock uint
var NumClock uint
var LivelloSblocco [2]int
var Possibilita [2]int
var ValoreNutrizionale int = 10
var ValoreNutrizionaleCarcassa int = 5
var Supestiti bool=true
//var ZonaCiboX int
//var ZonaCiboY int

func main() { //FUNZIONE MAIN
	SaluteIniziale [0]=50
	SaluteIniziale [1]=50
	CostoMovIniziale [0]=5
	CostoMovIniziale [1]=5
	CostoSexIniziale [0]=50
	CostoSexIniziale [1]=50
	EvoluzioneIniziale [0]=0
	EvoluzioneIniziale [1]=0
	AgeMaxInizio [0]=30
	AgeMaxInizio [1]=30
	LivelloSblocco [0]=0
	LivelloSblocco [1]=0
	Possibilita [0]=5
 	Possibilita [1]=5
	Clock = 1
	NumClock = 0
	rand.Seed(time.Now().UTC().UnixNano()) //inizializzazione rand
	/*height, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("height not valid")
	}
	width, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic("width not valid")
	}
	Altezza = height
	Larghezza = width*/
	cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()
	fmt.Println("\nInserisci altezza mondo [min = 10 | max = 20]: ")
	fmt.Scan(&Altezza)
	if Altezza > 20 {
		Altezza =20
	}
	if Altezza <10 {
		Altezza =10
	}
	fmt.Println("\nInserisci larghezza mondo [min = 10 | max = 20]: ")
	fmt.Scan(&Larghezza)
	if Larghezza > 20 {
		Larghezza =20
	}
	if Larghezza <10 {
		Larghezza =10
	}
	fmt.Println("\nInserisci 1 per usare i valori di default o un altro numero per medificarli: ")
	fmt.Scan(&def)
	if def!=1{
	fmt.Println("\nInserisci la salute iniziale: ")
	fmt.Println("Razza 1: ")
	fmt.Scan(&SaluteIniziale[0])
	fmt.Println("Razza 2: ")
	fmt.Scan(&SaluteIniziale[1])
	fmt.Println("\nInserisci il costo di uno spostamento iniziale (riduce l'energia ad ogni movimento) [default = 5]: ")
	fmt.Println("Razza 1: ")
	fmt.Scan(&CostoMovIniziale[0])
	fmt.Println("Razza 2: ")
	fmt.Scan(&CostoMovIniziale[1])
	fmt.Println("\nInserisci il costo di una riproduzione iniziale (riduce l'energia ad ogni riproduzione) [default = 50]: ")
	fmt.Println("Razza 1: ")
	fmt.Scan(&CostoSexIniziale[0])
	fmt.Println("Razza 2: ")
	fmt.Scan(&CostoSexIniziale[1])
	fmt.Println("\nInserisci i secondi di vita massimi (limita la durata della vita) [default = 30]: ")
	fmt.Println("Razza 1: ")
	fmt.Scan(&AgeMaxInizio[0])
	fmt.Println("Razza 2: ")
	fmt.Scan(&AgeMaxInizio[1])
	fmt.Println("\nInserisci il valore nutrizionale del cibo (di quanto aumenta l'energia di chi lo mangia) [default = 10]: ")
	fmt.Scan(&ValoreNutrizionale)
	fmt.Println("\nInserisci il valore nutrizionale delle carcasse (di quanto aumenta l'energia di chi lo mangia) [default = 5]: ")
	fmt.Scan(&ValoreNutrizionaleCarcassa)
	fmt.Println("\nInserisci il grado di evoluzione iniziale (se maggiore di zero migliora le prestazioni vitali)[default = 0]: ")
	fmt.Println("Razza 1: ")
	fmt.Scan(&EvoluzioneIniziale[0])
	fmt.Println("Razza 2: ")
	fmt.Scan(&EvoluzioneIniziale[1])
	fmt.Println("\nInserisci la possibilità di evoluzione (numero da 1 a 10) [default = 10]: ")
	fmt.Println("Razza 1: ")
	fmt.Scan(&Possibilita[0])
	if(Possibilita[0]<1 || Possibilita[0]>10) {
		Possibilita[0]=1}
		fmt.Println("Razza 2: ")
		fmt.Scan(&Possibilita[1])
	if(Possibilita[1]<1 || Possibilita[1]>10) {
		Possibilita[1]=1
	}
	fmt.Println("\nInserisci il livello di evoluzione visivo base (il livello evolutivo minimo che permette di vedere il cibo vicino)[default = 1]: ")
	fmt.Println("Razza 1: ")
	fmt.Scan(&LivelloSblocco[0])
	fmt.Println("Razza 2: ")
	fmt.Scan(&LivelloSblocco[1])
	}
	Matrix = make([][]*Element, Altezza)
	for i := range Matrix { // inizializzazione matrice
		Matrix[i] = make([]*Element, Larghezza)
		for j := range Matrix[i] {
			chose := rand.Intn(3) //scelta rando cibo bug o vuoto (null)
			switch chose {
			case 0:
				Matrix[i][j] = new(Element) // insetto
				Matrix[i][j].IsFood = false
				Matrix[i][j].Age = 0
				Matrix[i][j].Razza = rand.Intn(2)
				Matrix[i][j].Health = SaluteIniziale[Matrix[i][j].Razza]
				Matrix[i][j].CostoMov = CostoMovIniziale[Matrix[i][j].Razza]
				Matrix[i][j].CostoSex = CostoSexIniziale[Matrix[i][j].Razza]
				Matrix[i][j].Evoluzione = EvoluzioneIniziale[Matrix[i][j].Razza]
				Matrix[i][j].Premura = PremuraIniziale
				Matrix[i][j].AgeMax = AgeMaxInizio[Matrix[i][j].Razza]

			case 1:
				Matrix[i][j] = nil //vuota
			case 2:
				Matrix[i][j] = new(Element) // cibo
				Matrix[i][j].IsFood = true
				Matrix[i][j].Health = ValoreNutrizionale
				Matrix[i][j].Razza = 3
			}
		}
	}

	go ServiHTML() // fai partire il server html


	fmt.Println("Situazione iniziale: ")
	for Supestiti{
	aggiorna()
	rimasti()
	}
aggiorna()


}

func aggiorna() { //FUNZIONE AGGIORNA:	chiama la funzione muovi

		time.Sleep(time.Second * time.Duration(Clock))
		NumClock++
		for i := 0; i < Altezza; i++ {
			for j := 0; j < Larghezza; j++ {
					muovi(i, j)

			}
		}
		fmt.Printf("\nSituazione dopo %d movimenti:\n", NumClock)

}

func muovi(h int, w int) { //FUNZIONE MUOVI:	aggiorna la posizione di tutti gli oggetti in tabella	// h verticale, w orizzontale
	elemento := Matrix[h][w]                //assegnamente del contenuto della cella in 'elemento'
	if elemento == nil || elemento.IsFood { //controllo se 'elemento' è cibo o un altro essere
		return
	}



	if elemento.Health<=0 {
		fmt.Printf("Il soggetto in cella %d, %d è morto di fame\n",h, w)
		Matrix[h][w] = nil
		Matrix[h][w] = new(Element) // sostituisce con la carcassa
		Matrix[h][w].IsFood = true
		Matrix[h][w].Health = ValoreNutrizionaleCarcassa
		Matrix[h][w].Razza = 3
		return
	} else {
		elemento.Age++
	}

	if elemento.Age>Matrix[h][w].AgeMax {
		fmt.Printf("Il soggetto in cella %d, %d è morto di vecchiaia\n",h, w)
		Matrix[h][w] = nil
		Matrix[h][w] = new(Element) // sostituisce con la carcassa
		Matrix[h][w].IsFood = true
		Matrix[h][w].Health = ValoreNutrizionaleCarcassa
		Matrix[h][w].Razza = 3
		return
	}

	var h2 int
	var w2 int
	var direzCasOriz int
	var direzCasVert int
	var trovato bool = false

	if Matrix[h][w].Evoluzione>=LivelloSblocco[Matrix[h][w].Razza] {
	for i:=0;i<8;i++ {
		switch i {
	case 0:
		h2=-1
		w2=0
	case 1:
		h2=-1
		w2=1
	case 2:
		h2=0
		w2=1
	case 3:
		h2=1
		w2=1
	case 4:
		h2=1
		w2=0
	case 5:
		h2=1
		w2=-1
	case 6:
		h2=0
		w2=-1
	case 7:
		h2=-1
		w2=-1
		}
		if ((h+h2)>=0) && ((h+h2)<Altezza) && ((w+w2)>=0) && ((w+w2)<Larghezza) {
			if Matrix[h+h2][w+w2]!= nil {
				if Matrix[h+h2][w+w2].IsFood && !trovato {
				direzCasVert= h2;
				direzCasOriz = w2;
				trovato=true
				}
			}
		}
	}
}

	if !trovato {
		direzCasVert = rand.Intn(3)
		direzCasVert--
		direzCasOriz = rand.Intn(3)
		direzCasOriz--
	}
	nuovaPosizioneH := h + direzCasVert //aggiornamento posizione verticale
	nuovaPosizioneW := w + direzCasOriz //aggiornamento posizione orizzontale


	if nuovaPosizioneH >= Altezza || nuovaPosizioneH < 0 ||
		nuovaPosizioneW >= Larghezza || nuovaPosizioneW < 0 { //se esce dai bordi
		return
	}
	trovato=false
	if Matrix[nuovaPosizioneH][nuovaPosizioneW] != nil {

		if Matrix[nuovaPosizioneH][nuovaPosizioneW].Razza != Matrix[h][w].Razza { //se non è dalla stessa razza
			if Matrix[nuovaPosizioneH][nuovaPosizioneW].IsFood || (Matrix[nuovaPosizioneH][nuovaPosizioneW].Health+(Matrix[nuovaPosizioneH][nuovaPosizioneW].Evoluzione*5)) < (Matrix[h][w].Health+(Matrix[h][w].Evoluzione)*5) { // se e' cibo o un insetto piu debole
				Matrix[h][w].Health += Matrix[nuovaPosizioneH][nuovaPosizioneW].Health                //prelevamento energia essere fagocitato
				Matrix[nuovaPosizioneH][nuovaPosizioneW] = Matrix[h][w] //inglobamento essere perito
				Matrix[h][w] = nil
				Matrix[nuovaPosizioneH][nuovaPosizioneW].Health -= (Matrix[nuovaPosizioneH][nuovaPosizioneW].CostoMov)
				fmt.Printf("Il soggetto in cella %d, %d ha sconfitto quello in cella %d, %d\n",nuovaPosizioneH, nuovaPosizioneW, h, w)
			} else {	//perdita nel combattimento per la sopravvivenza
				Matrix[nuovaPosizioneH][nuovaPosizioneW].Health += Matrix[h][w].Health //il nemico prende l'energia
				Matrix[h][w] = nil
				fmt.Printf("Il soggetto in cella %d, %d ha fallito nel sconfiggere quello in cella %d, %d\n",h, w ,nuovaPosizioneH, nuovaPosizioneW)
			}
		} else { //se sono amici
			if nuovaPosizioneH == h && nuovaPosizioneW == w { //se cerca di mangiare il suo amico
				muovi(h, w)
			}
		}
	} else { //si muove sulla nuova casella
		Matrix[nuovaPosizioneH][nuovaPosizioneW] = Matrix[h][w]
		Matrix[nuovaPosizioneH][nuovaPosizioneW].Health -= Matrix[nuovaPosizioneH][nuovaPosizioneW].CostoMov
		Matrix[h][w] = nil

		if rand.Intn(Possibilita[Matrix[nuovaPosizioneH][nuovaPosizioneW].Razza]) == 0 { //se ha fortuna (o sfortuna) si evolve
			if rand.Intn(3) == 0 {
				Matrix[nuovaPosizioneH][nuovaPosizioneW].Evoluzione--
				Matrix[nuovaPosizioneH][nuovaPosizioneW].AgeMax-=5
			} else {
				Matrix[nuovaPosizioneH][nuovaPosizioneW].Evoluzione++
				Matrix[nuovaPosizioneH][nuovaPosizioneW].AgeMax+=10
			}
		}

		if (Matrix[nuovaPosizioneH][nuovaPosizioneW].Health-(Matrix[nuovaPosizioneH][nuovaPosizioneW].Premura)*10)>Matrix[nuovaPosizioneH][nuovaPosizioneW].CostoSex {		//se ha energia a sufficienza per riprodursi
			Matrix[h][w] = Costruttore(Matrix[nuovaPosizioneH][nuovaPosizioneW].Razza, Matrix[nuovaPosizioneH][nuovaPosizioneW].Evoluzione, Matrix[nuovaPosizioneH][nuovaPosizioneW].CostoMov, Matrix[nuovaPosizioneH][nuovaPosizioneW].CostoSex, Matrix[nuovaPosizioneH][nuovaPosizioneW].Premura, SaluteIniziale[Matrix[nuovaPosizioneH][nuovaPosizioneW].Razza], AgeMaxInizio[Matrix[nuovaPosizioneH][nuovaPosizioneW].Razza])
		}

	}

}
func rimasti(){
var i int
var j int
Supestiti=false
	for i=0;i<Altezza; i++{
		for j=0; j<Larghezza;j++{
			if Matrix[i][j]!=nil{
				if !Matrix[i][j].IsFood{
					Supestiti=true
				}
			}
		}
	}
}
/*
func giraMatrice(){
	var conta int=0
	var contaMax int=0
	var i int
	var j int
	for i=1;i<Altezza-1; i++{
		for j=1; j<Larghezza-1;j++{
			if Matrix[i][j]!=nil{
				if Matrix[i][j].IsFood{
					conta++;
				}
			}
			if Matrix[i-1][j]!=nil{
				if Matrix[i-1][j].IsFood{
					conta++;
				}
			}
			if Matrix[i-1][j-1]!=nil{
				if Matrix[i-1][j-1].IsFood{
					conta++;
				}
			}
			if Matrix[i-1][j+1]!=nil{
				if Matrix[i-1][j+1].IsFood{
					conta++;
				}
			}
			if Matrix[i][j+1]!=nil{
				if Matrix[i][j+1].IsFood{
					conta++;
				}
			}
			if Matrix[i][j-1]!=nil{
				if Matrix[i][j-1].IsFood{
					conta++;
				}
			}
			if Matrix[i+1][j-1]!=nil{
				if Matrix[i+1][j-1].IsFood{
					conta++;
				}
			}
			if Matrix[i+1][j]!=nil{
				if Matrix[i+1][j].IsFood{
					conta++;
				}
			}
			if Matrix[i+1][j+1]!=nil{
				if Matrix[i+1][j+1].IsFood{
					conta++;
				}
			}
				if conta>contaMax{
					ZonaCiboX=j
		 		 ZonaCiboY=i
				 contaMax=conta
				}
				conta=0;
		}
	}
	fmt.Printf("%d %d %d",contaMax,ZonaCiboX,ZonaCiboY)
}
*/

func stampaMatrice() {

	/*for i := 0; i < Altezza; i++ {
		for j := 0; j < Larghezza; j++ {
			if Matrix[i][j] == nil {
				fmt.Printf("   --  ")
			} else {
				if Matrix[i][j].IsFood {
					fmt.Printf("   CC  ")
				} else {
					fmt.Printf("%d   %d  ",Matrix[i][j].Razza, Matrix[i][j].Health)
				}
			}
		}
		fmt.Printf("\n")
	}*/
}
