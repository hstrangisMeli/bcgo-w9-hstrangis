package main

import "fmt"

type CustomError struct {
	Code int
	Msj  string
}

func NewCustomError(code int, msj string) *CustomError {
	return &CustomError{
		Code: code,
		Msj:  msj,
	}
}

func main() {
	salary := 50000
	if salary < 150000 {
		fmt.Printf("%v", NewCustomError(1, "Error: el salario ingresado no alcanza el mínimo imponible").Msj)
	}
}

/*Ejercicio 1 - Impuestos de salario #1
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario
ingresado no alcanza el mínimo imponible" y lanzalo en caso de que “salary” sea menor a 150.000. De lo contrario,
tendrás que imprimir por consola el mensaje “Debe pagar impuesto”.*/
