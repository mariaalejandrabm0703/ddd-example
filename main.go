package main

import "fmt"

type Hijo struct {
	edad int
}

func (h *Hijo) DameEdad() int {
	return h.edad
}

type Padre struct {
	hijo Hijo
}

func (p *Padre) DameEdadHijo() int {
	return p.hijo.DameEdad()
}

func main() {
	h := Hijo{edad: 4}
	fmt.Println("la edad de hijo es ", h.DameEdad())

	padre := Padre{hijo: h}
	fmt.Println("La edad del hijo es de: ", padre.DameEdadHijo())
}
