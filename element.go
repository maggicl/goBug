package main

import "fmt"

type Element struct { //struttura che contiene sia cibo sia amebe
	IsFood     bool			//se il contenuto della cella è cibo
	Health     int			//la sua vita
	Age        int			//la sua età
	AgeMax     int			//la sua età massima
	Razza      int			//per distiguere amici da nemici
	Evoluzione int			//se evolve in positivo avrà un bonus in attacco che viene sommato a Health
	CostoMov	 int			//quanta energia spende per muoversi
	CostoSex	 int			//quanto spende per riprodursi
	Premura	   int			//quanto distacco di energia è necessario per compiere la riproduzione allo scopo di evitare di rimanere a secco
}

func (e Element) String() string {
	return fmt.Sprintf("E'Cibo=%t Salute=%d Eta=%d", e.IsFood, e.Health, e.Age)
}

func Costruttore(razza int, evoluzione int, costomov int, costosex int, premura int, salute int, agemax int) *Element {
	nuovo := new(Element)
	nuovo.IsFood = false
	nuovo.Health = salute
	nuovo.Age = 0
	nuovo.Razza = razza
	nuovo.Evoluzione = evoluzione
	nuovo.CostoMov = costomov
	nuovo.CostoSex = costosex
	nuovo.Premura = premura
	nuovo.AgeMax = agemax
	return nuovo
}
