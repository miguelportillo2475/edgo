package main

import "fmt"

func main() {
	// Declaracion del if con la variable a nivel de función.
	var edad uint8 = 65
	if edad >= 66 {
		fmt.Println("Es un adulto de la tercera edad.")
	} else if edad >= 18 && edad <= 38 {
		fmt.Println("Es un adulto joven")
	} else if edad >= 39 && edad <= 65 {
		fmt.Println("Es un Adulto")
	} else {
		fmt.Println("Es un niño")
	}

	// Declaracion if con variable aisladad en el propio if
	if nombre := "Alejandro"; nombre == "Miguel" {
		fmt.Println("Hola Soy Miguel.")
	} else {
		fmt.Println("Hola, soy alguien más.")
	}
}
