// pacake main
// Es el paquete principal donde se practica este curso.
package main

// import  es la encargada de importar paquetes al proyecto.
import "fmt"

// func main es la funcion principal de inicio
func main() {
	//constante declarada pero no usada.
	const test int64 = 25

	// Contante en bloque
	// Deben ser declaradas y asignar si valor en la misma línea.
	const (
		nombre   string = "Miguel"
		apellido string = "Portillo"
	)

	//constantes en bloque, con tipado dinámico.
	const (
		segnombre   = "Alejandro"
		segapellido = "Boscán"
	)

	fmt.Println(nombre, segnombre, apellido, segapellido)

	// Usando el método IOTA para asignar valores en lista (iota vale 0)
	const (
		Enero = iota + 1
		Febrero
		Marzo
		Abril
		Mayo
		Junio
		Julio
		Agosto
		Septiembre
		Octubre
		Noviembre
		Diciembre
	)
	// imprimiendo las constantes de los meses.
	fmt.Println(Enero, Febrero, Marzo, Abril, Mayo, Junio, Julio, Agosto, Septiembre, Octubre, Noviembre, Diciembre)
}
