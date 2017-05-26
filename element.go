package main

import "fmt"

type Element struct { //struttura che contiene sia cibo sia amebe
	IsFood bool
	Health int
	Age    int
}

func (e Element) String() string {
	return fmt.Sprintf("<E'Cibo=%t Salute=%d Eta=%d>", e.IsFood, e.Health, e.Age)
}
