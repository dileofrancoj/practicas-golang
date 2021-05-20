package main

import "fmt"

/*
En la función main definir una variable de tipo "Producto".
Implementar dos funciones una que permita cargar el registro definido
en la main y otra que lo imprima.
*/

type Product struct {
	id			int
	description string
	price		float64
}

func load(product *Product) {
	fmt.Print("Ingrese el código: ")
	fmt.Scan(&product.id)
	fmt.Print("Ingrese la descripción: ")
	fmt.Scan(&product.description)
	fmt.Print("Ingrese el precio: ")
	fmt.Scan(&product.price)
}

func show(product Product) {
    fmt.Println("Codigo:", product.id)
    fmt.Println("Descripcion:", product.description)
    fmt.Println("Precio:", product.price)
}

func main(){
	var product1 Product
	load(&product1)
	show(product1)
}