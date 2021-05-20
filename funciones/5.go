package main

/*
Confeccionar una función que recibe un vector de 10 elementos enteros positivos,
retornar la cantidad de elementos con 1 dígito, 2 dígitos y más de dos dígitos.
*/

func digits(array [10]int) (oneDigit int, twoDigits int, threeOrMore int) {
	for i:=0; i < len(array); i++ {
		if array[i] < 10 {
			oneDigit++
		} else if(array[i] >= 10 && array[i] < 100) {
			twoDigits++
		} else {
			threeOrMore++
		}
	}
	return oneDigit,twoDigits,threeOrMore
}