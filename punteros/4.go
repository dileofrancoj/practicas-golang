package main

import "fmt"

/*
Elaborar una función que se le pase la dirección de una variable
entera e incremente en 1 lo apuntado por dicha variable.
*/

func increment(p *int) {
	*p = *p + 1
}

func main() {
	var v1 int = 4
	increment(&v1)
	fmt.Println(v1)
}