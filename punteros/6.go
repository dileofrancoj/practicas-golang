package main

import "fmt"

/*
Confeccionar un programa que contenga tres funciones:
1 - Cargar por teclado un vector de 5 elementos de tipo entero.
2 - Ordenar el vector de menor a mayor.
3 - Imprimir el vector. .
*/

func load(array *[5]int) {
	for i:=0; i < len(array); i++ {
		fmt.Print("Ingrese el elemento: ")
		fmt.Scan(&array[i])
	}
}

func order(array *[5]int) {
	for k:=0; k < len(array) - 1; k++ {
       for f := 0; f < len(array) - 1 - k; f++ {
			if array[f] > array[f+1] {
				aux := array[f]
				array[f] = array[f + 1]
				array[f+1]= aux
			}
		}
	}
}

func show(array [5]int){
	for i:=0;i<len(array);i++ {
		fmt.Print(array[i], " ")
	}
}

func main() {
	array := [5]int{2,4,5,6,1}
	load(&array)
	order(&array)
	show(array)
}