// package main es el inicio de mi programa de go
package main

// usando import en bloque
import (
	"fmt"
)

// todos lo operadores que no esten agrupados se comienzan con el nombre por ejemplo:
// Los primeros cuatro meses
const (
	enero = iota + 1
	febrero
	Marzo
	Abril
)

// main es la funcion inicial de mi programa
func main() {

}

// Funcion que contiene variables
func Variables() {

	// nombre Declaracion y asignacion de variables separadas.
	var nombre string
	nombre = "Miguel"

	// apellido Declaracion y asignación en la misma línea.
	var apellido string = "Portillo"

	// edad Declaracion y asignación de tipo Dinámica y con el método de variable corta (:=)
	edad := 29

	// Reasignación de Valores.
	nombre = "Miguel Alejandro"
	apellido = "Portillo Boscán"

	// var declaración en bloque.
	var (
		altura = 1.63
		peso   = 93
	)

	// Uso de las variables antes declaradas y asignadas en una impreson de consola.
	fmt.Println("Nombres: "+nombre, "Apellidos: "+apellido, "Altura: ", edad, altura, peso)

}

// Nota: Las constantes que no son grupos deben ser comentadas comenzando con el
// identificador de del la constante declarada.
// Ejem: const nombre = "Mark" = // nombre: Descripcion.

//Seccion de código comentado.
/*
func Constantes() {
	// test: constante declarada pero no usada.
	const test int64 = 25

	// Contante en bloque
	// Deben ser declaradas y asignar si valor en la misma línea.
	const (
		nombre   string = "Miguel"
		apellido string = "Portillo"
	)

	// constantes en bloque, con tipado dinámico.
	const (
		segnombre   = "Alejandro"
		segapellido = "Boscán"
	)
	fmt.Println(nombre, segnombre, apellido, segapellido)
}
*/
