package main

import (
	"fmt"
	"time"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    time.Time
}

type Alumnos struct {
	values []Alumno
}

var alumnos Alumnos = Alumnos{values: []Alumno{}}

func (alumnos *Alumnos) add(alumno Alumno) {
	alumnos.values = append(alumnos.values, alumno)
}

func (alumno Alumno) detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %s\nFecha: %s\n\n", alumno.Nombre, alumno.Apellido, alumno.DNI, alumno.Fecha)
}

func (alumnos Alumnos) Detalle() {
	for _, a := range alumnos.values {
		a.detalle()
	}
}
