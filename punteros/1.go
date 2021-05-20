package main

import "fmt"

/*
Definir dos variables enteras y almacenar valores por asignación.
Definir una variable puntero a entero
y guardar sucesivamente las direcciones de dichas dos variables y acceder a sus valores.
*/

func main() {
	var v1 int = 10
	var v2 int = 20
	var p *int
	p = &v1
	fmt.Println("Puntero a p: ", *p)
	fmt.Println("Dirección de p: ", p)
	p = &v2
	fmt.Println("Puntero a p: ", *p)
}