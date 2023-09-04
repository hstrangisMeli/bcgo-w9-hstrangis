package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
}

func (e CustomError) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible\n"
}

func ConsultarSalario(salary float64) (float64, error) {
	if salary < 10000 {
		return 0, CustomError{}
	}
	return salary, nil
}

func main() {
	_, err := ConsultarSalario(4000)
	if err != nil {
		if errors.Is(err, CustomError{}) {
			fmt.Println(err)
		}
	}
	return
}

/*Ejercicio 2 - Impuestos de salario #2
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje
“Error: el salario es menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000.
La validación debe ser hecha con la función “Is()” dentro del “main”.
*/
