package main

import (
	"fmt"
)

func main() {
	// Declaración Switch
	var ciclosDeEdad string = "Dead"
	switch ciclosDeEdad {
	case "Niño":
		fmt.Println("Es un niño.")
	case "Adolecente":
		fmt.Println("Es un adolecente.")
	// una manera de evaluar multiples casos.
	case "Adulto", "Adulto Joven":
		fmt.Println("Es un adulto.")
	case "Anciado":
		fmt.Println("Persona de la tercera edad.")
	default:
		fmt.Println("Es Difunto.")
	}

	// Switch Este se omite la empresion y se porne directamente en los case.
	persona := "Lex"
	switch {
	case persona == "Miguel" || persona == "Alejandro":
		fmt.Println("Hola Soy Miguel Portillo")
	default:
		fmt.Println("Hola, soy nadie.")
	}

	// Switch usando el operador lógico &&
	persona2 := "Portillo"
	edad := 30
	switch {
	case persona2 == "Miguel" || persona2 == "Alejandro":
		fmt.Println("Hola Soy Miguel Portillo")

	case persona2 == "Portillo" && edad == 29:
		fmt.Println("Hola, Soy Miguel Alejandro Portillo Boscan y tengo 29 Años.")
	default:
		fmt.Println("Hola.")
	}

	// Switch Usando el Operador de lógico de negación !
	var nombre string = "Miguel"
	var comprobar bool = false
	var interseccion bool = false
	switch {
	case !comprobar:
		fmt.Println("No se puede realizar la comprobación.")
	case nombre == "Daniel":
		fmt.Println("Soy Daniel")
	case nombre == "David":
		fmt.Println("Soy David.")
	case nombre == "Miguel" && interseccion == false:
		fmt.Println("Soy Miguel")
	default:
		fmt.Println("No soy Nadie")
	}
}