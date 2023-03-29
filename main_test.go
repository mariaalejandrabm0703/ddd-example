package main

import (
	"ddd-example/domain"
	"ddd-example/repository"
	"testing"
)

func TestNewLogService(t *testing.T) {
	hijo := repository.NewHijo(5)
	padre := domain.NewPadre(hijo)
	val := domain.NewValidationService(padre)
	log := domain.NewLogService(val)

	if log.DameEdadHijo() != 5 {
		t.Errorf("La edad del hijo en LogService debería ser 5, pero obtuve %d", log.DameEdadHijo())
	}
}

func TestNewValidationService(t *testing.T) {
	hijo := repository.NewHijo(5)
	padre := domain.NewPadre(hijo)
	val := domain.NewValidationService(padre)

	if val.DameEdadHijo() != 5 {
		t.Errorf("La edad del hijo en ValidationService debería ser 5, pero obtuve %d", val.DameEdadHijo())
	}
}

func TestNewPadre(t *testing.T) {
	hijo := repository.NewHijo(5)
	padre := domain.NewPadre(hijo)

	if padre.DameEdadHijo() != 5 {
		t.Errorf("La edad del hijo en Padre debería ser 5, pero obtuve %d", padre.DameEdadHijo())
	}
}

func TestNewHijo(t *testing.T) {
	hijo := repository.NewHijo(5)

	if hijo.DameEdad() != 5 {
		t.Errorf("La edad de Hijo debería ser 5, pero obtuve %d", hijo.DameEdad())
	}
}
