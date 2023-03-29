package repository

import "fmt"

type Hijo struct {
	edad int
}

func NewHijo(edad int) *Hijo {
	return &Hijo{
		edad: edad,
	}
}

func (h *Hijo) DameEdad() int {
	fmt.Println("retornando edad hijo")
	return h.edad
}
