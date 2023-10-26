package main

import "fmt"

// Panic y Recover

// Panic: nos permite detener nuestro programa.
// Recover: nos permite recobrarnos del panic y ejecutar el código o
// el resto de funciones para revisar otros potenciales errores.

func main() {

	division(100, 9)
	division(200, 10)
	division(500, 2)
	division(1, 0)
	division(99, 1)

}

// division recibe un dividendio y un divisor tipo entero, ambos.
// validad el divisor con otra función, de ser inválida, lanzará panic con msg entre paréntesis.
func division(dividend, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Me recuperé del panic")
		}
	}()
	validateZero(divisor)
	fmt.Println(dividend / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por cero")
	}
}

// Nota importante: la pila de llamada cuando hay un panic debemos
// leerla desde el final para dar con el error.

// pacnic: No se puede dividir por cero 	Mensaje del error
//
// goroutine 1 [running];
// main.validateZero(...)					Paquete main funcion validateZero archvio main.go, línea 26
// 		/.../.../.../main.go:26
// main.divison(...)						Paquete main funcion division
// 		/.../.../.../main.go:21
// main.main()     							Dentro de main en la fncion main
// 		/.../.../.../main.go:13 +0x25
// exit status 2

// usando con defer en una función anónima dentro y al inicio de la funcion division,
// ponemos una comprobación con el recover en caso de panic para ejecutar el resto de la función
// si queda algo por ejecutarse.
