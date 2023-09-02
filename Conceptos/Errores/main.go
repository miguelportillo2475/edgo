package main

import (
	"errors"
	"fmt"
	"strconv"
)

// se pueden comprobar errores de tipo específico.
var errNotFound = errors.New("Not Found")

// mapa para la funcion busqueda encadenada.
var food = map[int]string{
	1: "Hamburguesa",
	2: "Pizza",
	3: "Perro caliente",
}

func main() { ////////////////////////////////////////////////////////////////////////////////////////////////
	// Los errores son mejores que las exepciones porque son controlados por el usuario.
	// Los errores son parte importante de GO.
	// Lo ideal es que los errores se controlen apenas aparezcan, por eso las funciones devuelven
	// multiples valores.

	// La función Atoi del paquete strconv que trabaja con cadenas de texto
	// recibe un string y devuelve 2 valores, un entero y un error, y, para controlar
	// el error se usan 2 variables, la que contiene el valor y la del error que por convencion siempre se usa err

	// funcion que convierte valores cadenas de numeros en enteros y retorna  valor y error
	number, err := strconv.Atoi("5")
	if err != nil { // se comprueva si err es diferente a nil.
		fmt.Println(err) // se imprime el error si se encuentra.
		return           //se retorna porque la función falló.
	}
	fmt.Println(number) // se imprime el número  si noi hay error.
	fmt.Println("---------------------------------------------------")

	// ------------------------------------------------------------------
	//creamos y forzamos el error de la funcion search para ver el error específico.
	found, err := search("a") //obtenemos el valor y el error de search
	if err != nil {           // se comprueva si err es diferente a nil.
		fmt.Println(err) // se imprime el error si se encuentra.
		return           //se retorna porque la función falló.
	}
	fmt.Println(found) // imprimimos el valor

} ////////////////////////////////////////////////////////////////////////////////////////////////

// ------------------------------------------------------------------
// Podemos encadenar funciones para manejar los errores de multiples funciones.
// Funcion search encadenada
func search(key string) (string, error) {
	// recibe una llave de tipo string
	num, err := strconv.Atoi(key)
	if err != nil { // se comprueva si err es diferente a nil.
		return "", err //se retorna porque la función falló.
	}

	emoji, err := findFood(num)
	if err != nil { // se comprueva si err es diferente a nil.
		return "", err //se retorna porque la función falló.
	}
	return emoji, nil
}

func findFood(id int) (string, error) {
	value, ok := food[id]
	if !ok {
		return "", errNotFound
	}
	return value, nil
}
