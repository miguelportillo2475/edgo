// package main es el paquete principal del programa
package main

import "fmt"

// orden de prioridades de los operadores Aritméticos.
// 1:) () Entre paréntesis.
// 2:) *  Operador de Multiplicación
// 3:) /  Operador de División
// 4:) %  Operador de Módulo o porcentaje
// 5:) +  Operador de Suma
// 6:) -  Operador de resta

// main es la funcion primaria del aplicación
func main() {
	var a int8 = 25
	var b uint8 = 75
	c := 125
	//Operadores Aritméticos
	fmt.Println(2 + 3*3)                            // Valor: 11
	fmt.Println((2 + 3) * 3)                        // Valor: 15
	fmt.Println(float64(a) + float64(b)/float64(c)) // Valor 25.6  (Tríplecast)

	// Operadores de Asignación
	// 1:) =	Operador de igualdad
	var d = 25

	// 2:) +=	Operador para sumar otro valor al existente en la variable.
	d += 25

	// 3:) -=	Operador para restar otro valor al existente en la variable.
	d -= 12

	// 4:) *=	Operador para multiplicar otro valor con el valor de la variable.
	d *= 2

	// 5:) /=	Operador para dividir con otro valor con el de la variable.
	d /= 3

	// 6:) %=	Operador para Sacar el residuo a la variable con el valor pasado en la función.
	d %= 10

	fmt.Println("Valor Luego de 6 Operaciones:", d)

	// Declaraciones (No son operadores)
	// 1:) ++	Declaracion de post-incremento.
	d++
	fmt.Println("Valor Post-Incremento:", d)
	// 2:) --	Declaracion de post-decremento.
	d--
	fmt.Println("Valor Post-Decremento:", d)

	// Operadores de comparación (Booleanos mayormente dependiendo del resultado).

	e := 250
	f := 220
	g := 222

	// 1:) >	Mayor que...
	fmt.Println("Mayor(>):", e > g)

	// 2:) <	Menor que...
	fmt.Println("Menor(<):", f < e)

	// 3:) ==	Igual que...
	fmt.Println("Igual(==):", g == e)

	// 4:) !=	Diferente que...
	fmt.Println("Diferente(!=):", g != f)

	// 5:) >=	Mayor o igual que...
	fmt.Println("Mayo o Igual(>=):", f >= f)

	// 6:) <=	Menor o igual que...
	fmt.Println("Menor o Igual(<=):", e <= e)

	// Operadores Lógicos. ( &&=y / ||=o / !=onario o de negación )
	// edad Alamacena la edad del usuario.
	var edad byte = 29
	// Operador lógico ( y )
	fmt.Println("Es un adulto:", edad >= 18 && edad <= 60)

	// Operador lógico ( o )
	fmt.Println("Es niño o anciano:", edad < 18 || edad > 60)

	// Operador lógico Onario ( ! )
	fmt.Println("Onario para 4 mayor que 3:", !(4 > 3))
}
