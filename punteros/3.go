package main

import "fmt"

/*
Confeccionar una función que reciba como parámetros las direcciones de dos variables enteras y
le cargue a lo apuntado por dichas variables dos enteros.
*/

func load(p1 *int, p2 *int) {
	*p1 = 100
	*p2 = 150
}

func main() {
	var v1, v2 int
	load(&v1,&v2)
	fmt.Println(v1,v2)
}