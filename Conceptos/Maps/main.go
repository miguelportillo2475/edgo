// main paquete principal
package main

import (
	"fmt"
)

// func main Funcion principal
func main() {
	// Estructuras de Clave y Valor
	// music Mapa que almacena instrumentos musicales.
	music := make(map[string]string)
	music["ElectricGuitar"] = "Guitarra Eléctrica."
	music["AcusticGuitar"] = "Guitarra Acústica."
	music["Violin"] = "Violín."
	fmt.Println(music)

	// tech Mapa(con operador literal) que almacena productos de tecnología.
	tech := map[uint8]string{
		001: "Computador de mesa.",
		002: "Laptop.",
		003: "Tarjeta madre(Placa base).",
		004: "Mouse",
	}

	// delete funcion de eliminación.
	delete(tech, 01)
	fmt.Println(tech)

	// imprimir valor específico.
	fmt.Println(tech[2])

	// Imprimir valor no existente
	fmt.Println(tech[1])

	//content, ok Valorar si una clave existe.
	content, ok := tech[1]
	fmt.Println(content, ok)

}
