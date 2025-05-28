package clase2

import (
	"fmt"
	"strings"
)

func ContarLetras() {
	palabra := "Automovil"

	fmt.Println("Longitud: ", len(palabra))

	//Imprimir letra a letra

	splitPalabra := strings.SplitSeq(palabra, "")

	for value := range splitPalabra {
		fmt.Println(value)
	}
}

func Prestamos(edad int, empleado bool, sueldo float64, antiguedad int) {

	switch {
	case edad >= 22:
		fmt.Println("Cumple con la edad")
		fallthrough
	case empleado:
		fmt.Println("Esta empleado")
		fallthrough
	case sueldo >= 100000:
		fmt.Println("No se le aplica interes.")
		fallthrough
	case antiguedad > 1:
		fmt.Println("Cumple con todos lo criterios.")
	default:
		fmt.Println("No aplica para un prestamo.")
	}
}

func Mes(mes int) {

	if mes < 1 || mes > 12 {
		fmt.Println("El numero proporcionado no corresponde a un mes valido")
		return
	}
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	fmt.Println(meses[mes-1])
}

func Edad() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	mayor21anos := 0
	for nombre, value := range employees {
		if value > 21 {
			mayor21anos++
		}

		if nombre == "Benjamin" {
			fmt.Println("Edad de Benjamin: ", value)
		}

	}

	fmt.Println("Mayores de 21 años: ", mayor21anos)
	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Println(employees)
}
