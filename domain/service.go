package domain

import (
	"ddd-example/repository"
)

type Padre struct {
	hijo *repository.Hijo
}

func NewPadre(hijo *repository.Hijo) *Padre {
	return &Padre{
		hijo: hijo,
	}
}

func (p *Padre) DameEdadHijo() int {
	return p.hijo.DameEdad()
}
