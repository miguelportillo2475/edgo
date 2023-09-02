package main

import "fmt"

func main() {
	// colores Array normal
	var colores [5]string
	colores[0] = "Rojo"
	colores[1] = "Azul"
	colores[2] = "Amarillo"
	colores[3] = "Morado"
	colores[4] = "Verde"
	fmt.Println("Colores:", colores)

	// colores Array Literal sin limitar el indice
	colores1 := [...]string{"Rojo", "Azul", "Amarillo", "Morado", "Verde"}
	fmt.Println("Colores:", colores1)

}
