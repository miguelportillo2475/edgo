package main

// Varables variáticas y anonimas.
// siguiendo el ejemplo de la variable sums tenemos un problema
// al querer sumar más de un valor y aqui es donde entrans las funciones variáticas.

// Variáticas estas funciones convierten los parámetros en slices
// Pueden ser nombradas las funciones secundarias.
// paySum Función variática.
func paySum(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}
