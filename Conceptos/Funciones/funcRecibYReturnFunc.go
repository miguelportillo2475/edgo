package main

// Las funciones también son un tipo de dato, por lo que puedo usarlas como parametros
// y retornos de otras funciones.

// función como parametro.

// Función de filtrado del slice
func filter(nums []int, callback func(int) bool) []int {
	result := make([]int, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}
	return result
}

// Puedo Separar la parte de callback, asignarle nombre y ponerle distintas funciones
// greeterToTen Funcion mayor que diez
func greeterToTen(num int) bool {
	return num > 10
}

// lessThanTwenty Funcion menor que veinte
func lessThanTwenty(num int) bool {
	return num < 20
}

// Funciones con retorno de funciones.
// funcion que recibe como parámetro "a" y retorna una función "func sums(a int)"
// que recibe un entero y retorna un entero "func(int) int"
// esto sirve para reutilizar la funcion de la lógica (es para eso que se usa este ejemplo simple)
func sums(a int) func(int) int {
	// implemento la función en el return
	return func(b int) int {
		return a + b

	}
}
