package domain_test

import (
	"testing"

	"ddd-example/domain"
)

type ServiceMock struct{}

func (s *ServiceMock) DameEdadHijo() int {
	return 10
}

func TestDameEdadHijoL(t *testing.T) {
	mockService := &ServiceMock{}
	logService := domain.NewLogService(mockService)
	edad := logService.DameEdadHijo()
	if edad != 10 {
		t.Errorf("Expected edad to be 10. Got %d", edad)
	}
}
