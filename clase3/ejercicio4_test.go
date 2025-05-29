package clase3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImpuesto_impuesto(t *testing.T) {
	//Arrange
	expected := 40.77
	response := Inpuestos(151.000)
	//Act
	assert.Equal(t, expected, response)
}

func TestImpuesto_impuesto_fail(t *testing.T) {
	//Arrange
	expected := 9.350000000000001
	response := Inpuestos(55.000)
	//Act
	assert.Equal(t, expected, response)
}

func TestPromedio(t *testing.T) {

	expected := 3.2
	response, err := Promedio(2.4, 5, 3.5, 2.1, 3)

	assert.Equal(t, expected, response)
	assert.Nil(t, err)
	assert.NoError(t, err)
}

func TestPromedio_fail(t *testing.T) {
	_, err := Promedio(2.4, -5, 3.5, 2.1, 3)

	assert.Equal(t, err.Error(), "No se puede calcular el promedio con numeros negativos o superiores a 5")
	assert.NotNil(t, err)
	assert.Error(t, err)
}

// Test de Salarios
func TestCalcularSalario_categoriaA(t *testing.T) {

	expeted := "Total salario: 373.500"
	response, err := CalcularSalario(5000, "A")

	assert.Equal(t, expeted, response)
	assert.Nil(t, err)
	assert.NoError(t, err)
}

func TestCalcularSalario_categoriaB(t *testing.T) {

	expeted := "Total salario: 149.400"
	response, err := CalcularSalario(5000, "B")

	assert.Equal(t, expeted, response)
	assert.Nil(t, err)
	assert.NoError(t, err)
}
func TestCalcularSalario_categoriaC(t *testing.T) {

	expeted := "Total salario: 83.000"
	response, err := CalcularSalario(5000, "C")

	assert.Equal(t, expeted, response)
	assert.Nil(t, err)
	assert.NoError(t, err)
}

func TestCalcularSalario_fail(t *testing.T) {
	response, err := CalcularSalario(5000, "D")
	assert.Empty(t, response)
	assert.Equal(t, err.Error(), "La categoria no aplica.")
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestCalcularSalario_zeroMinutes(t *testing.T) {

	expeted := "Los minutos trabajados no corresponde a un valor valido."
	response, err := CalcularSalario(0, "A")

	assert.Empty(t, response)
	assert.Equal(t, err.Error(), expeted)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

// Test de estadisticas
func TestOperacion_minimun(t *testing.T) {
	expeted := 2
	response, err := Operacion("minimun")
	assert.Nil(t, err)
	resOperation, err := response(2, 3, 4, 5, 6, 7, 8, 9)
	assert.Equal(t, expeted, resOperation)
	assert.Nil(t, err)
}
func TestOperacion_minimun_fail(t *testing.T) {
	expeted := "El listado de numeros no puede estar vacio."
	response, err := Operacion("minimun")
	assert.Nil(t, err)
	resOperation, err := response()
	assert.Equal(t, 0, resOperation)
	assert.Equal(t, err.Error(), expeted)
	assert.NotNil(t, err)
}
func TestOperacion_maximun(t *testing.T) {
	expeted := 9
	response, err := Operacion("maximun")
	assert.Nil(t, err)
	resOperation, err := response(2, 3, 4, 5, 6, 7, 8, 9)
	assert.Equal(t, expeted, resOperation)
	assert.Nil(t, err)
}

func TestOperacion_average(t *testing.T) {
	expeted := 5
	response, err := Operacion("average")
	assert.Nil(t, err)
	resOperation, err := response(2, 3, 4, 5, 6, 7, 8, 9)
	assert.Equal(t, expeted, resOperation)
	assert.Nil(t, err)
}

func TestOperacion_nooperation(t *testing.T) {
	expeted := "Error seleccionando operación."
	response, err := Operacion("other")

	assert.Nil(t, response)
	assert.Equal(t, expeted, err.Error())
	assert.Error(t, err)
}

// Animales

func TestAnimal_notexist(t *testing.T) {
	expeted := "Animal no exist!"
	response, err := Animal("notexist")
	assert.Nil(t, response)
	assert.Equal(t, expeted, err.Error())
	assert.Error(t, err)
}

func TestAnimal_dog(t *testing.T) {
	var expeted float64 = 100
	response, err := Animal("dog")
	assert.Nil(t, err)
	assert.NoError(t, err)

	resDog, err := response(10)
	assert.Nil(t, err)
	assert.Equal(t, expeted, resDog)
}

func TestAnimal_cat(t *testing.T) {
	var expeted float64 = 50
	response, err := Animal("cat")
	assert.Nil(t, err)
	assert.NoError(t, err)

	resDog, err := response(10)
	assert.Nil(t, err)
	assert.Equal(t, expeted, resDog)
}
