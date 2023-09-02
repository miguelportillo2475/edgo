// main Paquete principal
package main

import "fmt"

func main() {
	// color Para optener la dirección en memoria se antespone un ampersant
	var color string = "rojo"
	fmt.Printf("Tipo: %T Valor:%s Direccion: %v\n", color, color, &color)

	// pointercolor Puntero al color y su valor.
	var pointercolor *string = &color
	fmt.Printf("Tipo: %T Valor:%v\n", pointercolor, pointercolor)

	// Cambiar el valor de una variable desde el apuntador
	*pointercolor = "Azul"

	// Referenciación
	fmt.Printf("Tipo:%T Valor:%v Referenciacion:%s\n", pointercolor, pointercolor, *pointercolor)

}
