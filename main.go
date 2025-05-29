package main

import (
	"fmt"

	"github.com/f-leo-rey/go-learning/clase3"
)

func main() {
	//clase1.Ejercicio1()
	//clase1.Ejercicio2()
	//clase2.ContarLetras()
	//clase2.Mes(24)
	//clase2.Mes(0)
	//clase2.Mes(4)
	//clase2.Mes(6)
	//clase2.Edad()
	//clase3.Ejercicio4()

	//Calcular Impuestos
	res1 := clase3.Inpuestos(151.000)
	res2 := clase3.Inpuestos(55.000)

	fmt.Println(fmt.Sprintf("$%.3f", res1))
	fmt.Println(fmt.Sprintf("$%.3f", res2))

	//Calcular promedio
	promedio, err := clase3.Promedio()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Promedio: ", promedio)

	//Calcular salario
	salarioA, err := clase3.CalcularSalario(5000, "A")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(salarioA)

	salarioB, err := clase3.CalcularSalario(5000, "B")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(salarioB)

	salarioC, err := clase3.CalcularSalario(5000, "C")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(salarioC)

	//Estadistica

	minimun, _ := clase3.Operacion("minimun")
	resMinimun, _ := minimun(9, 2, 3, 0, 5, 1)
	fmt.Println("Numero mas bajo: ", resMinimun)

	average, _ := clase3.Operacion("average")
	resAverage, _ := average(3, 5, 6, 7, 8, 9)
	fmt.Println("Promedio:", resAverage)

	maximun, _ := clase3.Operacion("maximun")
	resmaximun, _ := maximun(9, 2, 3, 0, 5, 1)
	fmt.Println("Numero mas alto: ", resmaximun)

	//Calcular cantidad alimento

	dog, err := clase3.Animal("dog")
	if err != nil {
		fmt.Println(err.Error())
	}
	resdog, _ := dog(10)
	fmt.Printf("Alimento requerido: %vKG\n", resdog)

	//cat
	cat, err := clase3.Animal("cat")
	if err != nil {
		fmt.Println(err.Error())
	}
	rescat, _ := cat(10)
	fmt.Printf("Alimento requerido: %vKG\n", rescat)

	hamster, err := clase3.Animal("hamster")
	if err != nil {
		fmt.Println(err.Error())
	}
	reshamster, _ := hamster(10)
	fmt.Printf("Alimento requerido: %vG\n", reshamster)

	tarantula, err := clase3.Animal("tarantula")
	if err != nil {
		fmt.Println(err.Error())
	}
	restarantula, _ := tarantula(10)
	fmt.Printf("Alimento requerido: %vG\n", restarantula)
}
