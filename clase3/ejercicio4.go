package clase3

import (
	"errors"
	"fmt"
)

func Inpuestos(salario float64) float64 {

	inpuesto1 := 0.17
	impuesto2 := 0.27
	if salario >= 50.000 && salario <= 150.000 {
		return salario * inpuesto1
	}

	if salario > 150.000 {
		return salario * impuesto2
	}

	return 0
}

func Promedio(notas ...float64) (float64, error) {
	suma := 0.0
	for _, valor := range notas {
		if valor < 0 || valor > 5 {
			return 0, errors.New("No se puede calcular el promedio con numeros negativos o superiores a 5")
		}
		suma += valor
	}
	return suma / float64(len(notas)), nil
}

func CalcularSalario(minutos int, categoria string) (string, error) {

	var (
		categoriaC float64 = 1.000
		categoriaB float64 = 1.500
		categoriaA float64 = 3.000
	)

	if minutos <= 0 {
		return "", errors.New("Los minutos trabajados no corresponde a un valor valido.")
	}

	horasTrabajadas := minutos / 60
	switch categoria {
	case "A":
		return total(horasTrabajadas, categoriaA, 0.50), nil
	case "B":
		return total(horasTrabajadas, categoriaB, 0.20), nil
	case "C":
		return total(horasTrabajadas, categoriaC, 0), nil
	default:
		return "", errors.New("La categoria no aplica.")
	}
}

func total(horasTrabajadas int, valorCategoria float64, extra float64) string {
	valorHoras := float64(horasTrabajadas) * valorCategoria
	valorExtra := float64(valorHoras) * extra
	total := float64(valorHoras) + valorExtra
	return fmt.Sprintf("Total salario: %.3f", total)
}

func Operacion(tipo string) (func(numeros ...int) (int, error), error) {

	switch tipo {
	case "minimun":
		return minimun, nil
	case "average":
		return average, nil
	case "maximun":
		return maximun, nil
	}

	return nil, errors.New("Error seleccionando operación.")
}

func minimun(numeros ...int) (int, error) {
	if len(numeros) == 0 {
		return 0, errors.New("El listado de numeros no puede estar vacio.")
	}

	min := numeros[0]

	for _, value := range numeros {
		if value < min {
			min = value
		}
	}

	return min, nil
}

func average(numeros ...int) (int, error) {

	if len(numeros) == 0 {
		return 0, errors.New("El listado de numeros no puede estar vacio.")
	}
	count := len(numeros)
	suma := 0
	for _, value := range numeros {
		suma += value
	}

	return suma / count, nil
}

func maximun(numeros ...int) (int, error) {
	if len(numeros) == 0 {
		return 0, errors.New("El listado de numeros no puede estar vacio.")
	}
	max := numeros[0]
	for _, value := range numeros {
		if value > max {
			max = value
		}
	}
	return max, nil
}

func Animal(typeAnimal string) (func(quantity float64) (float64, error), error) {
	switch typeAnimal {
	case "dog":
		return animalDog, nil
	case "cat":
		return animalCat, nil
	case "hamster":
		return animalHamter, nil
	case "tarantula":
		return animalTarantula, nil
	}

	return nil, errors.New("Animal no exist!")
}

func animalDog(quantity float64) (float64, error) {
	if quantity < 1 {
		return 0, errors.New("Quantity is invalid value.")
	}
	return (quantity * 10), nil
}

func animalCat(quantity float64) (float64, error) {
	if quantity < 1 {
		return 0, errors.New("Quantity is invalid value.")
	}
	return (quantity * 5), nil
}

func animalHamter(quantity float64) (float64, error) {
	if quantity < 1 {
		return 0, errors.New("Quantity is invalid value.")
	}
	return (quantity * 0.25), nil
}

func animalTarantula(quantity float64) (float64, error) {
	if quantity < 1 {
		return 0, errors.New("Quantity is invalid value.")
	}
	return (quantity * 0.15), nil
}
