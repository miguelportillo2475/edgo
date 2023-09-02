package main

import (
	"fmt"
	"os"
)

// la función defer: nos permite diferir/aplazar alguna tarea, se usa mucho a la hora
// de abrir y cerrar archivos, entre otras cosas, se aplazan las tareas hasta que la funcion
// desde donde fue llamada retorne.
// se pone mejor luego de la primera funcion de apertura y si se usan más con retornos, esta se ejecutará
// y cerrará las conexiones o archivos.
// con defer puedo: limpiar recursos, crear archivos, abri y cerrar conexión de red
// o crear controladores de bases de datos

// defer lista los aplazos en una lista ascendente es decir que el primer defer es el ultimo en ejecutarse
// y el ultimo será el primero: ejemplo 1

// el defer guarda la funcion ya ejecutada y esto hace que si los datos cambian, igualmente, entregará
// la versión antigua de los datos: ejemplo 2

//
//

func main() {
	// Ejemplo 1
	defer fmt.Println("primer defer - últumo en ejecutarse")
	defer fmt.Println("segundo defer - segundo en ejecutarse")
	defer fmt.Println("último defer - primero en ejecutarse")

	// Ejemplo 2
	num := 4
	defer fmt.Println("Valor1:", num)
	num = 10
	fmt.Println("Valor2:", num)

	// Ejemplo 3
	file, err := os.Create("test.txt") // creamos el archivo o da error
	if err != nil {                    //comprobamos el error
		fmt.Println(err) // imprimimos el error
		return           //retornams en caso de error
	}
	defer file.Close() //cerramos el archivos con defer para hacer el cierre apenas acabe la función
	defer fmt.Println("Cerrado")

	_, err = file.Write([]byte("Hello gophers")) //escribimos en el archivos [file.Write([]byte("")) crea un slice de bytes y
	if err != nil {                              // los va escribiendo en el archivos]
		fmt.Println(err) // imprimimos el error
	}

	// con defer nos ahorramos código redundante porque cada vez que trabajamos con un archivo
	// en una misma funcion o secuencia de ellas, entoces hay que cerrarlo en cada caso de retorno
	// pero con defer se ejecuta al finalizar y asi no quedan abiertos, y evitamos fugaz de memoria.
}
