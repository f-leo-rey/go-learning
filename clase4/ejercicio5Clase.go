package clase4

import (
	"encoding/json"
	"errors"
	"fmt"
)

/**************Product***************/
type Crud interface {
	Save() error
	GetAll() []Product
	GetById(id int) (Product, error)
}

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var products = []Product{
	{Id: 3, Name: "Monitor 27 curvo", Price: 120.450, Description: "Monitor curvo", Category: "Monitores"},
	{Id: 4, Name: "Mouse optico", Price: 80.450, Description: "Mouse Gamer", Category: "Gamer"},
	{Id: 5, Name: "Teclado ergonomico", Price: 56.789, Description: "Teclado ergonomico gamer", Category: "Gamer"},
	{Id: 6, Name: "Diadema", Price: 350.567, Description: "Diadema 365 gamer", Category: "Gamer"},
}

func (p Product) Save() error {
	//Validar que no exista
	for _, value := range products {
		if value.Id == p.Id {
			return errors.New("El identificador del producto ya existe.")
		}
	}
	products = append(products, p)
	return nil
}

func (p Product) GetAll() []Product {
	return products
}

// GetById implements Crud.
func (p Product) GetById(id int) (Product, error) {
	if id <= 0 {
		return Product{}, errors.New("El identificador no es valido")
	}
	var produc Product
	for _, value := range products {
		if value.Id == id {
			produc = value
			break
		}
	}
	return produc, nil
}

func Execute(c Crud, id int) {

	err := c.Save()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	productAll := c.GetAll()
	product1Json, err := json.MarshalIndent(productAll, "", " ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(product1Json))
	productId, err := c.GetById(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	product1GetJson, err := json.MarshalIndent(productId, "", " ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Recuperando registro por id: ", string(product1GetJson))
}

/**************Product***************/

/**************Employes***************/

type Person struct {
	Id          int
	Name        string
	DateOfBirth string
}

type Employe struct {
	Id       int
	Position string
	Person
}

func (e Employe) PrintEmployee() {
	fmt.Println(e)
}

/**************Employes***************/

func ManagementProduct() {
	product1 := Product{Id: 7, Name: "Pad Mouse", Description: "Pad Mouse 1x60cm", Price: 90.000, Category: "Gamer"}
	idProduc1 := product1.Id

	Execute(product1, idProduc1)

	product2 := Product{Id: 8, Name: "Protector de pantantalla", Description: "Protector de pantalla", Price: 20.000, Category: "Gamer"}
	idproduct2 := product2.Id

	Execute(product2, idproduct2)
}
