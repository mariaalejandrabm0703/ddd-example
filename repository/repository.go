package repository

type Hijo struct {
	edad int
}

func (h *Hijo) DameEdad() int {
	return h.edad
}
