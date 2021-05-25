package main

import "fmt"

/*
	Un struct puede tener asociado métodos para procesar los campos del struct.
	Un método es una función que tiene un parámetro extra de un determinado tipo de struct.

*/

/*
Declarar un struct que permita almacenar el nombre de una persona y su edad.
Definir los métodos que permitan:
1 Cargar el nombre y la edad de persona
2 Imprimir el nombre y la edad
3 Mostrar un mensaje si es mayor de edad
En la main definir dos variables del tipo struct declarado y llamar a sus métodos
*/

type Person struct {
	name string
	age  int
}

func (person *Person) setInfo() {
	fmt.Print("Ingrese el nombre de la persona: ")
	fmt.Scan(&person.name)
	fmt.Print("Ingrese la edad de la persona: ")
	fmt.Scan(&person.age)
}

func (person *Person) showInfo() {
	fmt.Println("Nombre: ", person.name)
	fmt.Println("Edad: ", person.age)
}

func (person *Person) isAdult() bool {
	if person.age >= 18 {
		return true
	}
	return false
}

func main() {
	var p1 Person
	p1.setInfo()
	p1.showInfo()
	isAdultp1 := p1.isAdult()
	fmt.Print(p1.name, " es adult@? ", isAdultp1)
}