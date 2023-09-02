package main

import "fmt"

// person Struct que almacena caracteristicas de una persona.
type person struct {
	Name      string
	LastName  string
	Age       uint8
	HavePets  bool
	PetsCount uint8
	PetsName  string
}

func main() {
	// Los Struct's permiten alamacenar una colección de campos, son similares a las clases de otros lenguajes.
	// miguel Instacia del Struct Person
	miguel := person{
		Name:      "Miguel",
		LastName:  "Portillo",
		Age:       29,
		HavePets:  true,
		PetsCount: 1,
		PetsName:  "Winky",
	}
	fmt.Printf("%+v\n", miguel)
	fmt.Printf("Edad de Miguel: %v\n", miguel.Age)

	// carolina Instancia del Struct Person
	carolina := person{"Carolina", "Pedrozo", 27, false, 0, ""}

	fmt.Printf("%+v\n", carolina)
	fmt.Printf("Carolina tiene mascota?: %+v\n", carolina.HavePets)

	// caro it's a Pointer to carolina instance
	caro := &carolina
	fmt.Println(caro)

	// 2 Ejemplos de referenciación para modificar las instancias desde los punteros.
	// Ejemplo 1
	(*caro).HavePets = true

	// Ejemplo 2
	caro.PetsCount = 1
	caro.PetsName = "Winky"

	fmt.Printf("%+v\n", caro)
	fmt.Printf("%+v\n", carolina)

}
