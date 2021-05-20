package main

import "fmt"

/*
Desarrollar dos funciones que reciban como parámetro el valor del lado de un cuadrado.
La primera debe calcular y mostrar la superficie y la segunda calcular y mostrar el perímetro.
En la main llamar a las funciones pasando los valores enteros comprendidos entre 10 y 20
*/

func perimeter(s int) {
	calculatedPerimeter := s * 4
	fmt.Println("El perimetro es: ", calculatedPerimeter)
}

func surface(s int) {
	calculatedSurface := s * s
	fmt.Println("El área es: ", calculatedSurface)

}

func main() {
	var side int
	fmt.Print("Ingrese un lado del cuadrado: ")
	fmt.Scan(&side)
	perimeter(side)
	surface(side)
}