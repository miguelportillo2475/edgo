package main

import (
	"fmt"
)

func main() {
	// La declaración For es la unica declaracion de ciclos en Go
	// La declaración normal se compone de 3 partes
	// Declaración de inicio
	i := 1
	// Condición que se evalúa antes de cada iteración.
	for i <= 5 {
		// Blocque de código a ejecutar
		fmt.Println("Valor de I:", i)
		// Declaración posterior, que se finaliza una iteración
		i++
	}
	fmt.Println("--------------------------------------")

	// Nota: la declaración de inicio puede compartirse entre ciclos
	// For Forever
	for {
		if int64(i) == 15 {
			break
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("--------------------------------------")

	// Los ciclos for son muy usados para iterar en los slices
	// Este ejmplo integra el slice en el propio for, solo es bueno
	// para ejemplos cortos o que no crezcan a futuro porque sino se complica.
	for i, v := range []string{"Miguel", "Alejandro", "Portillo", "Boscan"} {
		fmt.Println("Índice:", i, "Valor:", v)
	}
	fmt.Println("--------------------------------------")

	// numeros es un slice de numeros para iterar con un ciclo for
	// También pueden usarse operadores logicos, de comparacion y aritmeticos.
	numeros := []uint8{2, 4, 6, 8, 10}
	for i := range numeros {
		numeros[i] *= 2
	}
	fmt.Println(numeros)

	// Declaración for sobre mapas.
	// El siguiente ejemplo itera de forma aleatoria.
	fmt.Println("--------------------------------------")
	food := map[string]string{
		"Food Pasta":    "Pasta Boloñesa",
		"Food Arepa":    "Arepa",
		"Food Frijoles": "Frijoles en coco",
		"Food Arróz":    "Arróz",
	}
	for key, value := range food {
		fmt.Println("Llave:", key, "Valor:", value)
	}
	fmt.Println("--------------------------------------")

	// Se puede iterar sobre in string
	// El siguiente metodo dara como resultado el valor unicode de cada letra.
	for i, v := range "Miguel" {
		fmt.Println("Índice:", i, "Valor:", v)
	}

	fmt.Println("--------------------------------------")

	// Iterar sobre un string con casting sobre el valor para mostrar las letras
	// y no el unicode de cada letra de la cadena de caracteres.
	for i, v := range "Miguel" {
		fmt.Println("Índice:", i, "Valor:", string(v))
	}
}
