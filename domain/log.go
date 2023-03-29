package domain

import "fmt"

type log struct {
	s Service
}

func NewLogService(s Service) Service {
	return &log{
		s: s,
	}
}

func (l *log) DameEdadHijo() int {
	fmt.Println("mostrando unos logs")
	return l.s.DameEdadHijo()
}
