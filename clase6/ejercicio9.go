package clase6

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Employe struct {
	Name       string  `json:"name"`
	Salary     float64 `json:"salary"`
	TotalHours int     `json:"total_hours"`
}

type resultado struct {
	Date  string  `json:"date"`
	Total float64 `json:"total"`
}

var employes []Employe

func ReadJson() {
	os.Setenv("FILE_NAME", "employee_month_payment.json")
	os.Setenv("PATH_FILE", "/Users/frrey/Documents/")
	file := os.Getenv("FILE_NAME")
	path := os.Getenv("PATH_FILE")

	pathFile := path + file

	reader, err := os.ReadFile(pathFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal([]byte(reader), &employes)

	if err != nil {
		fmt.Println(err)
		return
	}

	total, err := SumarSalary(employes)
	if err != nil {
		fmt.Println(err)
	}

	result := resultado{Date: time.Now().Format("2006/01/02"), Total: total}

	fmt.Println(result)

	err = CrearJson(result)
	if err != nil {
		panic(err)
	}
}

func SumarSalary(employes []Employe) (float64, error) {
	if len(employes) == 0 {
		return 0, errors.New("Collection empty!")
	}
	total := 0.0
	for _, value := range employes {
		total += value.Salary
	}

	return total, nil
}

func CrearJson(res resultado) error {

	path := os.Getenv("PATH_FILE")
	jsonValue, err := json.Marshal(res)
	if err != nil {
		return errors.New("Error de conversión.")
	}

	filename := path + "totales-" + time.Now().Format("20060102") + ".json"

	err = os.WriteFile(filename, []byte(jsonValue), 0644)

	if err != nil {
		return errors.New("Error de escritura")
	}

	return nil
}
