package clase1

import (
	"fmt"
	"math"
)

const (
	nombre    = "Leonardo Rey"
	direccion = "Calle 10 # 5-32"
)

func Ejercicio1() {

	fmt.Println(nombre)
	fmt.Println(direccion)

	fmt.Println("###################################")

	var temperatura int = 18
	var humedad uint = 63
	var presionatmosferica float64 = 1018

	fmt.Printf(`
		Temperatura: %v°C 
		Presion: %vhPa
		Humedad: %v %%`,
		temperatura, humedad, presionatmosferica)

	//if
	x, n := 1.2, 9.4
	lim := 10.5

	if v := math.Pow(x, n); v < lim {
		fmt.Println("ASignación sobre if:", v)
	}

	numero1 := 30
	condicion := 20

	switch valor := condicion >= numero1; valor {
	case !valor:
		fmt.Println("La condición se cumple", valor)
	default:
		fmt.Println("No se cumple ninguna")
	}

	//Bucles

	//clasic for
	for i := 0; i <= 100; i++ {
		fmt.Println("Numero: ", i)
	}

	//while
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

	/*
		for {
			fmt.Println("for infinito")
		}*/

	frutas := []string{"Banano", "Manzana", "Papaya", "Pera"}

	for i, value := range frutas {
		fmt.Println("Indice: ", i, " Valor: ", value)
	}

	//continue validando pares e impares

	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i, "es inpar")
	}

	//Break
	for i := 0; i <= 20; i++ {
		if i >= 15 {
			break
		}

		fmt.Println(i)
	}

	//Ejercicios con bucles y switch

	paises := []string{"India", "Colombia", "Mexico", "Venezuela", "Paraguay"}

	for _, value := range paises {
		switch value {
		case "India":
			fmt.Println("El pais: ", value, " no pertenece a latam")
			fallthrough
		case "Colombia":
			fmt.Println(value, " esta en la lista")
		default:
			fmt.Println("Pasa algo")

		}
	}

	/*fmt.Println("###################################")
	var firstName string
	var lastName string
	var age int
	lastName := ""
	var driver_license = true
	var person string
	var height int
	childsNumber := 2

	fmt.Println("###################################")

	var lastName string = "Smith"
	var age int = 35
	boolean := false
	var salary float64 = 45857.90
	var firstName string = "Mary"*/
}
