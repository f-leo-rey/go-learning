package main

import (
	"github.com/f-leo-rey/go-learning/clase6"
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
	/*res1 := clase3.Inpuestos(151.000)
	res2 := clase3.Inpuestos(55.000)

	fmt.Println(res1)
	fmt.Println(res2)

	fmt.Println(fmt.Sprintf("$%.3f", res1))
	fmt.Println(fmt.Sprintf("$%.3f", res2))*/
	/*
		//Calcular promedio
		promedio, err := clase3.Promedio()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Promedio: ", promedio)
	*/
	//Calcular salario
	/*salarioA, err := clase3.CalcularSalario(5000, "A")

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

	fmt.Println(salarioC)*/

	//Estadistica
	/*
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
		fmt.Printf("Alimento requerido: %vG\n", restarantula)*/

	//Testing

	//Struct

	//clase4.PrintLibro()
	/*person := clase4.Person{Id: 1, Name: "Jose", DateOfBirth: "23-10-1992"}
	employe := clase4.Employe{Id: 20, Position: "Manager", Person: person}
	employe.PrintEmployee()

	clase4.ManagementProduct()*/

	//clase4.Show()

	/*clase4.Punteros()

	p1 := clase4.Stock{Name: "Moto", Cost: 50.000, Category: clase4.Category{Name: "Small", Percentage: 0}}
	var p2 *clase4.Stock
	p2 = &p1
	prod := p1.Factory(p2)
	fmt.Println(prod.Price())

	p3 := clase4.Stock{Name: "Moto", Cost: 50.000, Category: clase4.Category{Name: "Medium", Percentage: 1.3}}
	var p4 *clase4.Stock
	p4 = &p3
	prod2 := p3.Factory(p4)
	fmt.Println(prod2.Price())

	p5 := clase4.Stock{Name: "Moto", Cost: 50.000, ShippingCost: 2.500, Category: clase4.Category{Name: "Large", Percentage: 1.6}}
	var p6 *clase4.Stock
	p6 = &p5
	prod3 := p5.Factory(p6)
	fmt.Println(prod3.Price())*/
	/*err := clase5.Salary(100000)
	if err != nil {
		fmt.Println(err.Error())
	}

	clase5.Salary2(5000)
	clase5.Salary3(5000)
	clase5.Salary4(5000)

	total, err := clase5.CalculteSalary(80, 2000)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(total)*/

	//clase5.IsPair(2, 4, 6, 8, 7)

	/*clase5.ReadFile("customers.txt")
	fmt.Println("ejecución finalizada")

	clase5.ReadCustomers("customers.txtr")*/

	/*transaction := []clase5.Transaction{
		{From: clase5.Account{Nombre: "Jack", Saldo: 100_000}, To: clase5.Account{Nombre: "Doe", Saldo: 100_000}, Ammount: 50_000},
		{From: clase5.Account{Nombre: "Pedro", Saldo: 200_000}, To: clase5.Account{Nombre: "Juan", Saldo: 100_000}, Ammount: 400_000},
	}

	clase5.Transfer(&transaction)

	fmt.Println(transaction)*/

	clase6.ReadJson()
}
