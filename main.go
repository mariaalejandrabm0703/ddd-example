package main

import (
	"ddd-example/domain"
	"ddd-example/repository"
	"fmt"
)

func main() {
	hijo := repository.NewHijo(5)
	fmt.Println("Hijo- la edad de hijo es ", hijo.DameEdad())

	padre := domain.NewPadre(hijo)
	fmt.Println("Padre- La edad del hijo es: ", padre.DameEdadHijo())

	val := domain.NewValidationService(padre)
	fmt.Println("Validation- La edad del hijo es: ", val.DameEdadHijo())

	fmt.Println("////////////////////")

	log := domain.NewLogService(val)
	fmt.Println("Log- La edad del hijo es: ", log.DameEdadHijo())
}
