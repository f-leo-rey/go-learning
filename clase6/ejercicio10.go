package clase6

import (
	"encoding/csv"
	"errors"
	"os"
)

/*
Ejercicio 1: Inventario de productos (con archivo .txt)
Descripción:

	El archivo productos.txt contiene una lista de productos separados por coma:

	1,Mouse,20.5,10
	2,Teclado,45.0,5
	3,Monitor,150.0,2

Formato: ID,Nombre,Precio,Cantidad

Objetivo:

	Leer el archivo, cargar los productos en memoria (slice o map).

	Implementar funciones para buscar por ID, agregar stock, editar y eliminar stock.
		Si un producto tiene cantidad negativa, debe lanzar panic, y atraparlo con recover.

		Bonus:
		Usar

Usar defer para cerrar el archivo y/o capturar errores.

	Implementar Tests
*/
type Inventario struct {
	Id       string
	Nombre   string
	Precio   string
	Cantidad string
}

type Products []Inventario

func (p *Products) AddNew(i Inventario) {
	*p = append(*p, i)
}

func (p *Products) GetAll() []Inventario {
	return *p
}

func (p *Products) Edit(i Inventario) {

	for idx := range *p {
		if (*p)[idx].Id == i.Id {
			(*p)[idx].Nombre = i.Nombre
			(*p)[idx].Cantidad = i.Cantidad
			(*p)[idx].Precio = i.Precio
		}
	}
}

func (p *Products) Delete(id string) {
	var i int
	for ix := range *p {
		if (*p)[ix].Id == id {
			i = ix
		}
	}
	*p = append((*p)[:i], (*p)[i+1:]...)
}

func (p *Products) GetById(id string) (Inventario, error) {

	var inventario Inventario

	if id == "" {
		return inventario, errors.New("Identificador errado")
	}

	for _, value := range *p {
		if value.Id == id {
			inventario = value
		}
	}
	return inventario, nil
}

func WriterFile(data []Inventario) {
	os.Setenv("FILE_NAME", "productos.csv")
	os.Setenv("PATH_FILE", "/Users/frrey/Documents/")
	file := os.Getenv("FILE_NAME")
	path := os.Getenv("PATH_FILE")
	pathFile := path + file

	f, err := os.Create(pathFile)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	for _, value := range data {
		row := []string{
			value.Id,
			value.Nombre,
			value.Precio,
			value.Cantidad,
		}
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
}

func ReadFile() Products {
	file := os.Getenv("FILE_NAME")
	path := os.Getenv("PATH_FILE")
	pathFile := path + file

	var process []Inventario

	f, err := os.Open(pathFile)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for _, value := range records {
		inventario := &Inventario{
			Id:       value[0],
			Nombre:   value[1],
			Precio:   value[2],
			Cantidad: value[3],
		}

		process = append(process, *inventario)
	}
	return process
}
