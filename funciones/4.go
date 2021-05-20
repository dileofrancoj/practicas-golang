package main

import "fmt"

/*
Confeccionar una función que le enviemos un array de 5 enteros y nos retorne el mayor y menor en ese orden.
En la main definir y cargar un array.
*/

func higherLowerArray(array [5]int) (int,int){
	higher := array[0]
	lower := array[0]

	for i:=1; i < len(array); i++ {
		if array[i] > higher {
			higher = array[i]
		} else {
			if array[i] < lower {
				lower = array[i]
			}
		}
}
	return higher,lower
}

func main() {
	array := [5]int{2,4,5,6,1}
	h, l := higherLowerArray(array)
	fmt.Println("El mayor es:" ,h)
	fmt.Println("El menor es:" ,l)
}