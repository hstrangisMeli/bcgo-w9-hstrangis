package main

import "fmt"

var Products = []Product{}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, prod := range Products {
		fmt.Println(prod)
	}
}

func getById(id int) Product {
	for _, prod := range Products {
		if prod.ID == id {
			return prod
		}
	}
	return Product{}
}

func main() {
	p1 := Product{
		1,
		"Tomate",
		500,
		"Rojo",
		"Verdura",
	}
	p2 := Product{
		2,
		"Lechuga",
		450,
		"Criolla",
		"Verdura",
	}
	p3 := Product{
		3,
		"Manzana",
		730,
		"Verde",
		"Fruta",
	}
	p1.Save()
	p2.Save()
	p3.Save()
	p1.GetAll()
	fmt.Println(getById(3))
}

/*Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.
El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main().*/
