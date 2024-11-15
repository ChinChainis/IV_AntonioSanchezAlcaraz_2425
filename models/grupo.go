package models

import "errors"

type Turno string

const (
	Tarde  Turno = "Tarde"
	Maniana Turno = "Maniana"
)

type Grupo struct {
	Nombre string
	Turno  Turno
}

func NewGrupo(nombre string, turno Turno) (*Grupo, error) {
	
	if turno != Maniana && turno != Tarde {
		return nil, errors.New("turno inválido: debe ser 'Maniana' o 'Tarde'")
	}

	return &Grupo{
		Nombre: nombre,
		Turno:  turno,
	}, nil
}
