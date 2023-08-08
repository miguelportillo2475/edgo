package main

import "fmt"

func main() {

	// colores Array Literal sin limitar el índice.
	// Rojo[0] Azul[1] Amarillo[2] Morado[3] Verde[4]
	colores := [...]string{"Rojo", "Azul", "Amarillo", "Morado", "Verde"}
	fmt.Println("Colores:", colores)

	//primeros Slice Nota: los slice son de indice excluyente.
	primeros := colores[0:3] // Rojo[0] azul[1] amarillo[2]
	fmt.Println("Primeros 3 Colores:", primeros)

	// ultimos Slice de los ultimos colores.
	ultimos := colores[3:5] // Morado[3] verde[4]
	fmt.Println("Últimos 2 Colores:", ultimos)

	// Los datos de los slices se convierten en su índice.
	fmt.Println("Slice primeros:", primeros[1]) // Azul[1]
	fmt.Println("Slice ultimos:", ultimos[0])   // Morado[0]

	// ultimos Cambiemos el morado por blanco.
	ultimos[0] = "Blanco"

	// primeors Cambiemos el amarillos por violeta.
	primeros[2] = "Violeta"

	fmt.Println("Primeros 3 Colores:", primeros) // Rojo[0] azul[1] violeta[2]
	fmt.Println("Últimos 2 Colores:", ultimos)   // Blanco[3] verde[4]
	fmt.Println("Colores:", colores)

	// Inferir los índices
	// primeros2 Infiere desde el índice cero
	primeros2 := colores[:3] // desde el índice [0] al índice [2]
	fmt.Println(primeros2)

	// todos Infiere todo el índice.
	todos := colores[:]
	fmt.Println(todos)

	// Los Slices en go tienen Tamaño y Capacidad
	// len() cap()
	// animals Es un arreglo que animales.
	var animals = [5]string{"Perro", "Gato", "Loro", "Elefante", "Jirafa"}
	fmt.Println(animals)

	// pets Slice  para imprimir las mascotas
	pets := animals[:2]
	fmt.Println("Mascotas:", pets)
	fmt.Println("Tamaño:", len(pets))
	fmt.Println("Capacidad:", cap(pets))

	// pets Agregar con append nuevas mascotas y verificar tamaño y capacidad de nuevo
	pets = append(pets, "Canario", "MiniPig")
	fmt.Println("Nuevas Mascotas:", pets)
	fmt.Println("Tamaño:", len(pets))
	fmt.Println("Capacidad:", cap(pets))

	// Inicializa un Slice sin referenciar un array y esto creara un array interno en memoria.
	// food Slice sin referencia a un array
	food := []string{"Tomate", "Cebolla", "Pan", "Cereales", "Res", "Cerdo", "Pollo"}
	fmt.Println(food)
	fmt.Println(len(food), "Productos de:", cap(food), "Productos.")
	fmt.Printf("Tipo: %T Valor: %v\n", food, food)

	// Creando un Slice con la funcion (Make)
	// Variable corta / Make(Cadena, tamaño, Capacidad)
	fruits := make([]string, 0, 3)
	fruits = append(fruits, "apple", "orange", "pear")
	fmt.Println("Tamaño:", len(fruits), "Capacidad:", cap(fruits))
	fmt.Printf("Tipo: %T Valor: %v\n", fruits, fruits)

	// Comprobando el valor cero o nulo (nil)
	fmt.Println("Valor Cero fruits:", fruits == nil)

	// Los valores inicializados (Con llaves) dan como resultado (false)
	valorCerofalse := []string{}
	fmt.Println("Valor Cero Inicializado:", valorCerofalse == nil)

	// valorCerotrue Comparando con nil para ver el valor cero
	var valorCerotrue []string
	fmt.Println("Valor Cero Verdadero:", valorCerotrue == nil)

}
