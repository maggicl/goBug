package main

import "fmt"

type Element struct { //struttura che contiene sia cibo sia amebe
	IsFood     bool			//se il contenuto della cella è cibo
	Health     int			//la sua vita
	Age        int			//la sua età
	Razza      string		//per distiguere amici da nemici
	Evoluzione int			//se evolve in positivo avrà un bonus in attacco che viene sommato a Health
	CostoMov	 int			//quanta energia spende per muoversi
	CostoSex	 int			//quanto spende per riprodursi
}

func (e Element) String() string {
	return fmt.Sprintf("<E'Cibo=%t Salute=%d Eta=%d>", e.IsFood, e.Health, e.Age)
}
