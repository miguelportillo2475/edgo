package main

import "fmt"

func main() {
	//constante declarada pero no usada.
	const test int64 = 25

	// Contante en bloque
	// Deben ser declaradas y asignar si valor en la misma línea.
	const (
		nombre   string = "Miguel"
		apellido string = "Portillo"
	)

	//constantes en bloque, con tipado dinámico.
	const (
		segnombre   = "Alejandro"
		segapellido = "Boscán"
	)
	fmt.Println(nombre, segnombre, apellido, segapellido)
}
