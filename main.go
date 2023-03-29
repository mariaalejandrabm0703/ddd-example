package main

import (
	"ddd-example/repository"
	"fmt"
)

type Padre struct {
	hijo *repository.Hijo
}

func (p *Padre) DameEdadHijo() int {
	return p.hijo.DameEdad()
}

func main() {
	h := repository.NewHijo(5)
	fmt.Println("Hijo- la edad de hijo es ", h.DameEdad())

	padre := Padre{hijo: h}
	fmt.Println("Padre- La edad del hijo es: ", padre.DameEdadHijo())
}
