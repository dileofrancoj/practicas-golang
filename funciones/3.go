package main

import "fmt"

/*

Confeccionar una función que le enviemos dos enteros y nos retorne el mayor y menor en ese orden.
En la main solicitar la carga de dos valores por teclado y llamar a la función.

*/

func higherLower(n1,n2 int) (int, int){
	if n1 > n2 {
		return n1,n2
	}
	return n2 ,n1
}

func main() {
	var n1, n2 int
	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&n1)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&n2)
	h, l := higherLower(n1,n2)
	fmt.Println("El mayor es: ", h)
	fmt.Println("El menor es: ", l)
}