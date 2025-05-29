package clase3

import "fmt"

func Ejercicio4() {
	resultado := suma(3, 3)
	fmt.Println(resultado)

	cuadrado := Cuadrado(suma(8, 9))

	fmt.Println(cuadrado)

	fmt.Print("Uso de funciones con elipsis")

	Elipsis(2, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6)

	fmt.Print("Uso de funciones que retornan funciones\n")

	opeSuma := orquestador("Sum")

	res := opeSuma(2, 5)

	fmt.Println("Resultado de la suma es: ", res)

	opeResta := orquestador("resta")
	resre := opeResta(4, 6)
	fmt.Println("Resultado de la resta es: ", +resre)

}

// Funciones con retorno nombrado
// Al tener la variable de retorno nombrada no es necesario  hacer el retorno explicitamente de la misma
// Esto genera un error
func suma(num1 int, num2 int) (resultado int) {
	resultado = num1 * num2
	return
}

// Como no se nombro la variable de retorno si la debo retonar de forma explicita
func Cuadrado(num int) int {
	resultado := num * num
	return resultado
}

// Multiples valores por parametro del mismo tipo
func Elipsis(valores ...int) {
	for _, value := range valores {
		value += value
		fmt.Println(value)
	}
}

//funciones que retornan funciones meszclado con retornos nombrados

func orquestador(tipo string) func(num1 int, num2 int) int {

	switch tipo {
	case "Sum":
		return sumas
	case "resta":
		return resta
	}
	return nil
}

func sumas(num1 int, num2 int) (resultado int) {
	resultado = num1 + num2
	return
}

func resta(num1 int, num2 int) int {
	resultado := num1 - num2
	return resultado
}
