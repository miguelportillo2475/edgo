// package main inicio del paquete principal.
package main

import "fmt"

// func main es la funcion inicial del programa.
func main() {
	// booleana es una variable booleana(valor verdadero o falso)
	var booleana bool = false
	fmt.Printf("Varlor: %v Tipo de datos:%T\n", booleana, booleana)

	// UINT: solo almacena datos positivos:
	// uint8: 8 bit integer (0 a 255)
	// uint16: 16 bit integer (0 a 65535)
	// uint32: 32 bit integer (0 a 4294967295)
	// uint64: 64 bit integer (0 a 18446744073709551615)

	// INT: Permite almacenar datos numericos negativos aunque no se limita a ellos:
	// int8: 8 bit integer (-128 a 127)
	// int16: 16 bit integer (-32768 a 32768)
	// int32: 32 bit integer (-2147483648 a 2147483648)
	// int64: 64 bit integer (-9223372036854775808 a 9223372036854775808)

	// N: Almacenar un valor superior es overflows(desbordamineto)
	// Es importante saber para que se unsará una variable porque
	// asignar solo el tipo int  o uint usará los valores máximos,
	// lo que repercute en el espacio usado en memoria y el rendimiento.

	// byte: es el alias de uint8 (0 a 255)
	// rune: es el alias de int32 o unicode

	// edad Alamacena la edad con byte(uint8) como tipo de dato
	var edad byte
	edad = 29
	fmt.Printf("Tipo de datos: %T Valor: %v\n", edad, edad)

	// unic alamacena un valor unicode con el tipo de dato rune(int32)
	var unic rune = 'a'
	fmt.Printf("Tipo de datos: %T Valor: %v\n", unic, unic)

	// float32 Para valores de comas flotantes o decimales de 32 bits (533620.3)
	// float64 Para valores de comas flotantes o decimales de 64 bits (533620.316)

	// Casting de datos para operar entre tipos de datos numericos.
	//a variable tipo uint8
	var a uint8 = 255
	// b variable  tipo uint16
	var b uint16 = 2536

	/* error1 es un ejemplo de lo que no hay que hacer para castear variableas numericas
	error1 := a + b
	print(error1)
	*/

	// c casting de las variables a y b para operar con sus valores sin problemas de tipo de datos.
	c := uint16(a) + b
	fmt.Printf("Tipos de datos: %T Valor: %v\n", c, c)

	// decimal variable para nueva pruba de casting a parte del curso
	var decimal float64 = 530000.316

	// entera es una variable para el casting con un decimal
	var entera int = 3620

	// fusion es una prueba de casting de datos flotantes y enteros
	fusion := decimal + float64(entera)
	fmt.Println(fusion)

	// Variable sin uso  pero manteniendola sin comentarla para su posterior uso(ID Blank)
	var _ int16 = 925

	// Valor Cero de cada Variable
	// abc variable tipo string con valor cero
	var abc string
	fmt.Printf("tipo: %T Varlor: %q\n", abc, abc)

	// def variable numerica con valor cero
	var def int
	fmt.Printf("tipo: %T Valor: %v\n", def, def)

	// ghi variable numerica con valor cero
	var ghi uint
	fmt.Printf("tipo: %T Valor: %v\n", ghi, ghi)

	// jkl variable booleana con valor cero
	var jkl bool
	fmt.Printf("tipo: %T Valor: %v\n", jkl, jkl)
}
