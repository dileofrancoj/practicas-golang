package main

import "fmt"

/*
Declarar un registro que permita almacenar el codigo, descripcion y precio de un producto.
Luego definir dos variables de dicho tipo, cargarlas por teclado e imprimir el
nombre del producto que tiene mayor precio.
*/

type Producto struct {
	id 				int
	description 	string
	price 			float64
}

func main() {
	var p1, p2 Producto
	fmt.Print("Ingrese el código: ")
	fmt.Scan(&p1.id)
	fmt.Print("Ingrese la descripción: ")
	fmt.Scan(&p1.description)
	fmt.Print("Ingrese el precio: ")
	fmt.Scan(&p1.price)
	fmt.Println("****************************")
	fmt.Print("Ingrese el código: ")
	fmt.Scan(&p2.id)
	fmt.Print("Ingrese la descripción: ")
	fmt.Scan(&p2.description)
	fmt.Print("Ingrese el precio: ")
	fmt.Scan(&p2.price)

	if p1.price > p2.price {
		fmt.Print("El producto ", p1.id, " tiene un precio mayor")
	} else {
		fmt.Print("El producto ", p2.id, " tiene un precio mayor")	
	}
}