package main

import "fmt"

type Product struct{
	code 		int
	description int
	price	 	float64
}

// Definir un vector de 4 elementos de tipo Producto
// Cargar vector
// Imprimir vector
// Mostrar articulo de > precio

func load(products *[4]Product) {
	for f:=0; f < len(products); f++ {
		fmt.Print("Ingrese el código único del producto: ")
		fmt.Scan(&products[f].code)
		fmt.Print("Ingrese la descripción del producto: ")
		fmt.Scan(&products[f].description)
		fmt.Print("Ingrese el precio producto: ")
		fmt.Scan(&products[f].price)
	}
}

func show(products [4]Product){
	fmt.Println("Listado de productos")
	for f:=0; f<len(products);f++ {
		fmt.Println("Código: ", products[f].code)
		fmt.Println("Descripción: ", products[f].description)
		fmt.Println("Precio: ", products[f].price)
	}
}

func highPrice(products [4]Product) int {
	pos := 0
	for f:=1; f<len(products);f++ {
		if products[f].price > products[pos].price {
			pos = f
		}
	}
	return products[pos].code
}

func main() {
	var products [4]Product
	load(&products)
	show(products)
	highPriceProduct := highPrice(products)
	fmt.Print("El producto de mayor precio es el: ", highPriceProduct)
}