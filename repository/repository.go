package repository

type Hijo struct {
	edad int
}

func NewHijo(edad int) *Hijo {
	return &Hijo{
		edad: edad,
	}
}

func (h *Hijo) DameEdad() int {
	return h.edad
}
