// main asigna el inicio del paquete.
package main

// "fmt" paquete importado para formatear el texto e imprimir texto en consola.
import "fmt"

// main es la clase inicial.
func main() {
	// nombre Variable declarada y asignado el valor en una línea diferente.
	var nombre string
	nombre = "Miguel"

	// apellido Declara la variable y asigna valor en la misma línea.
	var apellido string = "Portillo"

	// edad Declara y asigna valor en la primera línea con el método de variable corta (:=).
	edad := 29

	// nombre Reasignando el nombre.
	nombre = "Miguel Alejandro"

	// apellido Reasignando el apellido.
	apellido = "Portillo Boscán"

	// Declaración en bloque.
	var (
		altura = 1.63
		peso   = 93
	)

	// Uso de las variables antes declaradas y asignadas en una impreson de consola.
	fmt.Println("Nombres: "+nombre, "Apellidos: "+apellido, "Altura: ", edad, altura, peso)

}
