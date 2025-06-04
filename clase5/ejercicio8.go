package clase5

import (
	"fmt"
	"io"
	"os"
)

func IsPair(num ...int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for _, value := range num {
		if (value % 2) != 0 {
			panic("Not an event number")
		}
		fmt.Println(value, " is an event number! ")
	}

}

/*******Ejercicio 1 ************/

func ReadFile(path string) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)

		}
	}()

	_, err := os.ReadFile(path)

	if err != nil {
		panic("he indicated file was not found or is damaged")
	}

}

/*******Ejercicio 2 ************/
/*
Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos
leer este archivo e imprimir los datos que contenga. En el caso de no poder leerlo, se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución,
independientemente del resultado, debemos tener un “defer” que nos indique que la ejecución finalizó.
También recordemos cerrar los archivos al finalizar su uso.
*/
func ReadCustomers(path string) {

	pathFile := "/Users/frrey/Documents/" + path

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.Open(pathFile)

	if err != nil {
		panic("Error de lectura!")
	}

	defer file.Close()
	content, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
	fmt.Println("ejecución finalizada")

}

/*******Ejercicio 3 ************/
/*
objetivo

Simular una serie de transferencias bancarias entre cuentas. El sistema debe:

Lanzar un panic si una cuenta intenta transferir más dinero del que tiene.

Usar defer para registrar las operaciones (log).
Usar recover para evitar que el programa se detenga si una operación falla.

	Requisitos
Define una estructura Account con nombre y saldo.

Crea una función Transfer(from *Account, to *Account, amount float64) que:

Verifique si from tiene fondos suficientes.
Lance panic si no tiene fondos.

Si tiene fondos, descuente y transfiera el dinero.

Usa defer para registrar la operación (aunque falle).

Usa recover para capturar el error y que el programa continúe con las siguientes operaciones.
*/
type Transaction struct {
	From    Account
	To      Account
	Ammount float64
}
type Account struct {
	Nombre string
	Saldo  float64
}

func Transfer(tran *[]Transaction) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, value := range *tran {
			value.From.Saldo -= value.Ammount
			value.To.Saldo += value.Ammount

			fmt.Printf(`Transfiriendo de la cuenta %v a la cuenta %v un total de %.2f
			`, value.From.Nombre, value.To.Nombre, value.Ammount)
		}

	}()

	for _, value := range *tran {
		if value.From.Saldo < value.Ammount {
			panic("Saldo insuficiente")

		}
	}

}
