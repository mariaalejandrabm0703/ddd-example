package main

import (
	son "ddd-example/repository/repository"
	"fmt"
)

type Padre struct {
	hijo son.Hijo
}

func (p *Padre) DameEdadHijo() int {
	return p.hijo.DameEdad()
}

func main() {
	h := son.Hijo{edad: 4}
	fmt.Println("la edad de hijo es ", h.DameEdad())

	padre := Padre{hijo: h}
	fmt.Println("La edad del hijo es de: ", padre.DameEdadHijo())
}
