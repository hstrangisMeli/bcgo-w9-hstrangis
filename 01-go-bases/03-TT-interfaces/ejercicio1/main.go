package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func NewAlumno(nombre string, apellido string, dni int, fecha string) (a *Alumno) {
	a = &Alumno{
		Nombre:   nombre,
		Apellido: apellido,
		DNI:      dni,
		Fecha:    fecha,
	}
	return
}

func (a Alumno) Detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	a1 := NewAlumno("Federico", "Gomez", 25666789, "01/09/2023")
	a1.Detalle()
}

/*Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle*/
