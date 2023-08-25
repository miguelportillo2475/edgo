package main

import "strings"

// Las funciones puende retornar valores pero deben ser declarados los tipos de valores que retornaran.
func sum(a, b int) int {
	return a / b
}

// las funciones pueden retornar varios valores a la vez.
func convert(text string) (string, string) {
	lower := strings.ToLower(text)
	upper := strings.ToUpper(text)
	return lower, upper

}

// se pueden nombrar los valores al igual que los parámetros:
// pero esto es recomendable solo cuando las funciones son cortas y no
// se prevee que crescan a futuro.
// ademas pueden ser agrupadas con sus tipos como los parámetros.
func convertion(text string) (nlower string, nupper string) {
	nlower = strings.ToLower(text)
	nupper = strings.ToUpper(text)
	return nlower, nupper

}
