package domain

type Padre struct {
	hijo Hijo
}

type Hijo interface {
	DameEdad() int
}

func NewPadre(hijo Hijo) *Padre {
	return &Padre{
		hijo: hijo,
	}
}

func (p *Padre) DameEdadHijo() int {
	return p.hijo.DameEdad()
}

type Service interface {
	DameEdadHijo() int
}
