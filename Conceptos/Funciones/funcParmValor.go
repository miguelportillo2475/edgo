package main

import (
	"fmt"
	"strings"
)

// Funciones con parámetros por valor y referencia.
// las funciones pueden ser llamadas desde otras funciones

func greet() {
	fmt.Println("Hola gopher")
}

// Función con Parámetros de nombre y apellido
func goodbye(nombre string, apellido string) {
	fmt.Println("Adiós", nombre, apellido)
}

// función para saludar de vuelta
// Se pueden agrupar parámetros por tipo de dato.
func sayHelloGoBack(nombre, apellido, estatus string, edad uint8) {
	fmt.Println("Hola, soy", nombre, apellido, estatus, edad, "años de edad.")
}

// Función con error, de referencia lógica:
// no funciona como se espera porque las funciones reciben una copia del valor.
/* func toUpperCaseErr(text string) {
	text = Strings.ToUpper()
}
*/
// Para poder hacerlo correctamente se deben usar punteros:
func toUpperCase(text *string) {
	*text = strings.ToUpper(*text)

}
