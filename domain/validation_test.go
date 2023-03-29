package domain_test

import (
	"ddd-example/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockService struct{}

func (s *mockService) DameEdadHijo() int {
	return 10
}

func TestNewValidationService(t *testing.T) {
	s := &mockService{}
	v := domain.NewValidationService(s)
	assert.NotNil(t, v, "Validation service should not be nill")
}

func TestDameEdadHijoV(t *testing.T) {
	s := &mockService{}
	v := domain.NewValidationService(s)
	edad := v.DameEdadHijo()
	assert.Equal(t, 10, edad, "Edad should be 10")
}
