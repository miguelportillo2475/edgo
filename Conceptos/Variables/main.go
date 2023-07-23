package main

import "fmt"

func main() {
	// Declaracion y asignacion de variables separadas.
	var nombre string
	nombre = "Miguel"

	// Declaracion y asignación en la misma línea.
	var apellido string = "Portillo"

	// Declaracion y asignación de tipo Dinámica y con el método de variable corta (:=)
	edad := 29

	// Reasignación de Valores.
	nombre = "Miguel Alejandro"
	apellido = "Portillo Boscán"

	//declaración en bloque.
	var (
		altura = 1.63
		peso   = 93
	)

	//Uso de las variables antes declaradas y asignadas en una impreson de consola.
	fmt.Println("Nombres: "+nombre, "Apellidos: "+apellido, "Altura: ", edad, altura, peso)

}
