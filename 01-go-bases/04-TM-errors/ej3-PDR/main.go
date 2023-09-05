package main

import (
	"errors"
	"fmt"
)

type Cliente struct {
	Legajo    int
	Nombre    string
	DNI       int
	Telefono  int
	Domicilio string
}

var Clientes = []Cliente{}

var (
	ErrDatosNoValidos = errors.New("Faltan datos para poder ingresar el cliente")
	ErrExisteCliente  = errors.New("Error: el cliente ya existe")
	ErrMultiples      = errors.New("Se detectaron varios errores en tiempo de ejecución")
)

func DatosNoValidos(cliente Cliente) (bool, error) {
	if cliente.Legajo == 0 || cliente.Nombre == "" || cliente.DNI == 0 || cliente.Telefono == 0 || cliente.Domicilio == "" {
		var errorDatosInvalidos = ErrDatosNoValidos
		return true, errorDatosInvalidos
	}
	return false, nil
}

func ExisteCliente(c Cliente) (bool, error) {
	for _, cliente := range Clientes {
		if c.Legajo == cliente.Legajo {
			var errorExisteCliente = ErrExisteCliente
			return true, errorExisteCliente
		}
	}
	return false, nil
}

func AgregarCliente(cliente Cliente) {
	var datosInvalidos, existeCliente bool
	var errDatosInvalidos, errExisteCliente error
	datosInvalidos, errDatosInvalidos = DatosNoValidos(cliente)
	existeCliente, errExisteCliente = ExisteCliente(cliente)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	switch {
	case datosInvalidos && existeCliente:
		panic(ErrMultiples)
	case errors.Is(errDatosInvalidos, ErrDatosNoValidos):
		panic(ErrDatosNoValidos)
	case errors.Is(errExisteCliente, ErrExisteCliente):
		panic(ErrExisteCliente)
	default:
		Clientes = append(Clientes, cliente)
	}
}

func main() {
	var c1 = Cliente{
		1,
		"Carlos",
		22656888,
		11556463,
		"Las Heras 1123",
	}
	var c2 = Cliente{
		1,
		"Carlos",
		0,
		11556463,
		"Las Heras 1123",
	}
	//Error de cliente repetido
	AgregarCliente(c1)
	AgregarCliente(c1)
	Clientes = Clientes[:0]

	//Error de datos faltantes
	AgregarCliente(c2)
	Clientes = Clientes[:0]

	//Ambos errores
	AgregarCliente(c1)
	AgregarCliente(c1)
	AgregarCliente(c2)

	fmt.Println("Fin de la ejecución")
}

/*Ejercicio 3 - Registrando clientes
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos son:
Legajo
Nombre
DNI
Número de teléfono
Domicilio


Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás leer los datos de un array. En caso de que esté repetido, debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “Error: el cliente ya existe”, y continuar con la ejecución del programa normalmente.
Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar que todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de ellos tendrá que ser del tipo error para el caso de que se ingrese por parámetro algún valor cero (recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).
Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes: “Fin de la ejecución” y “Se detectaron varios errores en tiempo de ejecución”. Utilizá defer para cumplir con este requerimiento.

Requerimientos generales:
Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente para el caso de error retornado).*/
