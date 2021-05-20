package main

import "fmt"

/*
	Confeccionar una función que reciba dos enteros e imprima el mayor de ellos.
	Llamar a la función desde la main cargando previamente dos valores por teclado.
*/

func higher(n1 , n2 int) {
	var mayor int
	if n1 > n2 { 
		mayor = n1 
	} else {
		mayor = n2
	}
	fmt.Println("El mayor es: " , mayor)
}

func main() {
	var n1, n2 int
	fmt.Print("Ingresá el primer número: ")
	fmt.Scan(&n1)
	fmt.Print("Ingresá el segundo número: ")
	fmt.Scan(&n2)
	higher(n1,n2)
}