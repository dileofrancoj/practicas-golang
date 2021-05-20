package main

import "fmt"

func main() {
    var s1 string = "uno"
    var s2 string ="dos"
    var ps *string
    ps = &s1
    fmt.Println(s1) //se imprime: uno
    *ps = "tres"
    fmt.Println(s1) //se imprime: tres
    s1 = "cuatro"
    fmt.Println(*ps) //se imprime: cuatro
    ps = &s2
    fmt.Println(*ps) //se imprime: dos
	// Al modificar de s1 o ps obtenemos el mismo resultado
}

