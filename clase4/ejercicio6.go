package clase4

import (
	"fmt"
	"time"
)

type Studend struct {
	Name     string
	Apellido string
	Dni      string
	Fecha    time.Time
}

func (s Studend) getNombre() string {
	return s.Name
}

func (s Studend) getApellido() string {
	return s.Apellido
}

func (s Studend) getDni() string {
	return s.Dni
}

func (s Studend) getFechaIngreso() time.Time {
	return s.Fecha
}

func Punteros() {

	studen := Studend{Name: "Carl", Apellido: "Jhonson", Dni: "123456789", Fecha: time.Now().Local()}

	fmt.Println(studen.getNombre())
	fmt.Println(studen.getApellido())
	fmt.Println(studen.getDni())
	fmt.Println(studen.getFechaIngreso())
}

type Product1 interface {
	Price() string
}

type Stock struct {
	Name         string
	Cost         float64
	ShippingCost float64
	Category
}

type Category struct {
	Name       string
	Percentage float64
}

func (s Stock) Price() string {
	return fmt.Sprintf("Costo con porcentaje incluido: %.3f", s.Cost)
}

func (s *Stock) Factory(s1 *Stock) Product1 {
	if s1.Category.Percentage != 0 {
		s.Cost = s1.Cost * s1.Category.Percentage
		s.Cost += s1.ShippingCost
	}
	return s
}
