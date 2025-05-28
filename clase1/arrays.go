package clase1

import "fmt"

func Ejercicio2() {
	//array

	var arrayNombre [2]string

	arrayNombre[0] = "Leonardo"
	arrayNombre[1] = "Freyder"

	fmt.Println(arrayNombre[0], arrayNombre[1])
	fmt.Println(arrayNombre)

	//slice tamaño dinamico

	var frutas = []string{"Mango", "Banano"}

	fmt.Println(frutas[0]) //Mango

	//slice definiendo longitud y capacidad

	sli := make([]string, 3, 5)

	fmt.Println(cap(sli))
	fmt.Println(len(sli))

	//Slice con make

	s := []string{}

	s = append(s, "prueba1")
	s = append(s, "Prueba2")
	s = append(s, "Prueba3")

	fmt.Println(s)

	//Eliminar elementos de un slice
	fmt.Println("Eliminando un elemento del slice")
	s = s[:len(s)-1]

	fmt.Println(s)

	fmt.Println("Eliminar un indice especifico")
	indice := 0
	s = append(s[:indice], s[:indice+1]...)
	fmt.Println(s)
	//Imprimir un rango
	fmt.Println("Imprimiendo un rango")
	fmt.Println(s[1:2]) //imprime prueba2 y prueba3

	//longitud y capacidad

	fmt.Println("Usando len ", len(s))
	fmt.Println("Usando cap ", cap(s))

	//Maps

	mapas := map[string]int{"a": 1, "b": 2, "c": 3}

	for key, value := range mapas {
		res := fmt.Sprintf("Key: %v Valor: %v ", key, value)
		fmt.Println(res)
	}

	//logitud de un map
	fmt.Println(len(mapas))

	//Acceder a elementos concretos
	fmt.Println(mapas["a"])

	//Agregar elementos al mapa
	mapas["d"] = 4

	fmt.Println(mapas)

	fmt.Println("Eliminar un elemento de un mapa")
	delete(mapas, "d")
	fmt.Println(mapas)
}
