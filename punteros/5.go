package main

import "fmt"

/*
Definir un vector de 5 elementos de tipo entero en la función main.
Implementar una función que retorne el mayor y el menor elemento del vector
por medio de dos parámetros de tipo puntero:
*/

func higherLower(array [5]int, h *int, l *int) {
	*h = array[0]
	*l = array[0]
	for i:=0;i<len(array);i++{
		if array[i] > *h {
			*h = array[i]
		} else {
			if array[i] < *l {
				*l = array[i]
			}
		}
	}
}

func main() {
	array:= [5]int{1,2,3,5,6}
	var h,l int
	higherLower(array, &h, &l)
	fmt.Println("Mayor: ", h)
	fmt.Println("Menor: ", l)

}