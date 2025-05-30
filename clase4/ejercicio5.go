package clase4

import (
	"fmt"
	"math"
)

type Login struct {
	User     string `json:"user"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Password string `json:"-"` //esta etiqueta permite omitir un campo privado opción en este caso para no mostrar la contraseña
}

type Libro struct {
	Isbn             string `json:"isbn,omitempty"` //Eliminar campos vacios
	Copias           int
	Titulo           string
	Descripcion      string
	FechaPublicacion string
	Paginas          int
	Autor            Autor
	Categoria        Categoria
}

type Categoria struct {
	Genero      string
	Descripcion string
}

type Autor struct {
	Nombre       string
	Apellido     string
	Nacionalidad string
}

/*****************Ejemplo de metodos**************/

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

/*****************Ejemplo de metodos**************/

func PrintLibro() {
	libro := Libro{
		Isbn:             "978-628-7669-42-0",
		Copias:           10_000,
		Titulo:           "Nexus",
		Descripcion:      "Una breve historia de las redes de información desde la edad de piedra hasta la IA",
		FechaPublicacion: "10 de septiembre de 2024",
		Paginas:          608,
		Autor: Autor{
			Nombre:       "Yubal Noah",
			Apellido:     "Harari",
			Nacionalidad: "israelí",
		},
		Categoria: Categoria{
			Genero:      "Debate",
			Descripcion: "Historia de las ideas",
		},
	}

	fmt.Println(libro)

	circle := Circle{radius: 4.6}

	area := circle.area()
	fmt.Println("Area del circulo: ", area)

	perimeter := circle.perimeter()
	fmt.Println("Perimetro del circulo: ", perimeter)
}
