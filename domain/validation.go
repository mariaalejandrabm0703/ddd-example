package domain

import "fmt"

type validation struct {
	s Service
}

func NewValidationService(s Service) Service {
	return &validation{
		s: s,
	}
}

func (l *validation) DameEdadHijo() int {
	fmt.Println("mostrando unos validations")
	return l.s.DameEdadHijo()
}
