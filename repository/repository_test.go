package repository_test

import (
	"testing"

	"ddd-example/repository"
)

func TestNewHijo(t *testing.T) {
	edad := 10
	hijo := repository.NewHijo(edad)

	if hijo == nil {
		t.Error("NewHijo debería retornar una instancia de Hijo, pero es nil")
	}
}

func TestDameEdad(t *testing.T) {
	edad := 10
	hijo := repository.NewHijo(edad)

	resultado := hijo.DameEdad()

	if resultado != edad {
		t.Errorf("DameEdad debería retornar %d, pero retornó %d", edad, resultado)
	}
}
