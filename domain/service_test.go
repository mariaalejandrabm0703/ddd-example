package domain_test

import (
	"ddd-example/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHijo struct{}

func (h *MockHijo) DameEdad() int {
	return 10
}

func TestPadre_DameEdadHijo(t *testing.T) {
	mockHijo := &MockHijo{}

	padre := domain.NewPadre(mockHijo)
	edad := padre.DameEdadHijo()

	// Verificar que se obtenga el valor esperado
	assert.Equal(t, 10, edad)
}
