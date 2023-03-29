package main

import (
	"ddd-example/domain"
	"ddd-example/repository"
	"fmt"
)

func main() {
	h := repository.NewHijo(5)
	fmt.Println("Hijo- la edad de hijo es ", h.DameEdad())

	padre := domain.NewPadre(h)
	fmt.Println("Padre- La edad del hijo es: ", padre.DameEdadHijo())
}
