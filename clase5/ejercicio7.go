package clase5

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

/**********Primer ejercicio***********/
type CustomErr struct {
	Code int
	Msg  string
}

func (c CustomErr) Error() string {
	return "error: the salary entered does not reach the taxable minimum"
}

func Salary(salary int) error {
	var error CustomErr
	if salary < 150000 {
		fmt.Println(error.Error())
	}
	return nil
}

/**********Segundo ejercicio***********/
type CustomErr2 struct {
	Code int
	Msg  string
}

var mycustomerror = new(CustomErr2)

func (c CustomErr2) Error() string {
	return fmt.Sprintf("error: salary is less than 10000")
}

func getSalary(salary int) error {
	if salary <= 10000 {
		return mycustomerror
	}
	return nil
}

func Salary2(salary int) {
	err := getSalary(salary)

	if errors.Is(err, mycustomerror) {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Great!")
}

/**********Tercer ejercicio***********/

func getSalary3(salary int) error {
	if salary <= 10000 {
		return errors.New("error: salary is less than 10000")
	}
	return nil
}

func Salary3(salary int) {
	err := getSalary3(salary)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Great!")
}

/**********Cuarto ejercicio***********/
func getSalary4(salary int) error {
	if salary <= 150000 {
		return fmt.Errorf("error: the minimum taxable amount is 150,000 and the salary entered is: %v", salary)
	}
	return nil
}

func Salary4(salary int) {
	err := getSalary4(salary)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Great!")
}

/**********Quinto ejercicio***********/

func CalculteSalary(hours int, valuexhours float64) (float64, error) {

	if hours < 0 || hours < 80 {
		return 0, fmt.Errorf("Error: the worker cannot have worked less than 80 hours per month")
	}

	total := float64(hours) * valuexhours

	if total > 150_000 {
		total = total - (total * 0.1)
	}
	return total, nil
}

func Errores() {
	//Unwrap
	error_1 := fmt.Errorf("Ocurrio un error")
	error_2 := fmt.Errorf("Ocurrio otro error %w", error_1)

	err := errors.Unwrap(error_2) //comprueba si la cadena tiene otro error y retorna ese error por eso el resultado es error_1

	if err == error_1 {
		fmt.Println("Ocurrio un error desconocido")
	}

	//IS
	//Valida si el tipo de error producido al intentar abrir el archivo es de tipo fs.ErrExist si coincide retorna true
	_, err = os.Open("archivo")
	var noexist error = fs.ErrNotExist
	if errors.Is(err, noexist) {
		fmt.Println("El archivo no existe")
	}

	//AS
	//Valida si el tipo de error producido al intentar abrir el archivo es de tipo *fs.PathError si coincide retorna true
	//En este caso se pasa el puntero
	//Esta función se comporta como un assertion de tipos

}
