package main

import (
	"fmt"
)

func main() {
	// Funciones con parámetros por valor y referencias
	greet() // llamando a la función saludar sin parárametros
	goodbye("Miguel", "Portillo")
	sayHelloGoBack("Miguel", "Portillo", "Gopher novato de casi", 30)
	fmt.Println("-------------------END 1-----------------------")
	// Funciones con parámetros por valor y referencias

	// Funciones con retorno.
	var name string = "Miguel Portillo" //Variable para el ejemplo To Upper y ToLower
	toUpperCase(&name)
	fmt.Println(name)
	// funciones que retornan valores
	result := sum(100, 10)
	fmt.Println(result)

	// Funciones con multiples valores
	lower, upper := convert("Miguel Alejandro Portillo Boscán.")
	fmt.Println("LowerCase:", lower, "UpperCase:", upper)

	// Funciones con multiples valores
	nlower, nupper := convertion("Gopher de corazón.")
	fmt.Println("LowerCase:", nlower, "UpperCase:", nupper)
	fmt.Println("------------------------------------------")
	// Funciones con retorno.

	// Funciones que reciven y devuelven Funciones.
	// slice  de numeros enteros
	nums := []int{2, 12, 25, 9, 18, 66}
	resp := filter(nums, func(num int) bool { return num > 10 })
	fmt.Println(resp)
	fmt.Println("--------------------------------------------")

	// De esta forma las funciones son más fáciles de usar y más limpias, además multifunciones.
	resp = filter(nums, greeterToTen)
	fmt.Println("Mayor que 10:", resp)
	resp = filter(nums, lessThanTwenty)
	fmt.Println("Menor que 20", resp)
	fmt.Println("--------------------------------------------")

	// Funciones con retorno de funciones.
	results := sums(2)(15)
	fmt.Println("Ejemplo de de funcion que retorna función:", results)

	// Funciones con retorno de funcciones de verdad utiles.
	plus10 := sums(10) // primer valor de la funcion sums

	fmt.Println("ejemplo:", plus10(25)) // impresion de plus10 + 2do valor de la función sums
	fmt.Println("ejemplo:", plus10(15))
	fmt.Println("ejemplo:", plus10(5))
	fmt.Println("ejemplo:", plus10(1))
	fmt.Println("ejemplo:", plus10(2))
	fmt.Println("--------------------------------------------")

	// La funciones Variáticas y anónimas.
	// Variáticas.

	sum1 := paySum(1)
	sum2 := paySum(2, 3)
	sum3 := paySum(15, 26, 93)
	sum4 := paySum(65, 65, 65, 65, 6659)
	sum5 := paySum(65, 5466, 65, 65, 6659, 65535)

	fmt.Println("suma de 1 valor:", sum1)
	fmt.Println("suma de 2 valores:", sum2)
	fmt.Println("suma de 3 valores:", sum3)
	fmt.Println("suma de 5 valores:", sum4)
	fmt.Println("suma de 6 valores:", sum5)
	fmt.Println("--------------------------------------------")

	// las funciones anónimas son iguales que las normales solo que no se les da nombre.
	// greet2 es una variable que contiene una función anónima.
	greet2 := func(name2 string) {
		fmt.Println("hola", name2)

	}
	greet2("Miguel")

	// este tipo de función se llama por el nombre de la variable que lo
	// contiene y se le pasan los parámetros

	// las funciones anónimas también pueden tener paramentros y ser auto ejecutadas sin
	// meterlas en una variable
	func(nombre string) {
		fmt.Println("Hola,", nombre)
	}("Gophers")

}
