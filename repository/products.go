package repository

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/product-api/data"
)

func GetAllProducts() data.ProductsList {

	db := DatabaseConnect()
	selectAll, err := db.Query("select * from product p")
	if err != nil {
		panic(err.Error())
	}

	product := &data.Product{}
	productList := data.ProductsList{}

	for selectAll.Next() {
		var id int
		var name, description string
		var price float32

		err = selectAll.Scan(&id, &name, &description, &price)
		if err != nil {
			panic(err.Error())
		}
		product.Name = name
		product.Description = description
		product.Price = price
		product.ID = id

		// change append
		productList.AddProduct(*product)
	}

	defer db.Close()
	return productList
}

func InsertNewProduct(r *http.Request) data.ProductsList {

	db := DatabaseConnect()
	prodFromJSON := data.Product{}
	prodFromJSON.FromJson(r.Body)
	insert, err := db.Prepare("insert into product (name, description, price) values ($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(prodFromJSON.Name, prodFromJSON.Description, prodFromJSON.Price)

	productList := data.ProductsList{}
	productList = GetAllProducts()
	defer db.Close()
	return productList
}

func DeleteProduct(r *http.Request) data.ProductsList {

	// melhorar o metodo de extraçao do parametro ID
	p := strings.Split(r.URL.Path, "/")

	db := DatabaseConnect()
	delete, err := db.Prepare("delete from product where product.id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(p[2])

	productList := data.ProductsList{}
	productList = GetAllProducts()
	defer db.Close()
	return productList
}

func UpdateProduct(r *http.Request) data.ProductsList {

	db := DatabaseConnect()
	prodFromJSON := data.Product{}

	productID := strings.Split(r.URL.Path, "/")

	prodFromJSON.FromJson(r.Body)
	insert, err := db.Prepare("update product set name = $1, description = $2, price = $3 where product.id = $4")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(prodFromJSON.Name, prodFromJSON.Description, prodFromJSON.Price, productID[2])

	productList := data.ProductsList{}
	productList = GetAllProducts()
	defer db.Close()
	return productList
}

func InitialData() data.ProductsList {

	connectionInfo := "user=postgres password=password host=localhost sslmode=disable"

	dbLocalConnection, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		panic(err.Error())
	}
	dbLocalConnection.Exec("create database products")

	defer dbLocalConnection.Close()

	db := DatabaseConnect()

	db.Exec("create table product (id serial primary key, name varchar, description varchar, price decimal)")

	db.Exec("insert into product (name, description, price) values ('tea', 'cup of tea', 7)")
	db.Exec("insert into product (name, description, price) values ('milk', 'box of milk', 5)")

	productList := data.ProductsList{}
	productList = GetAllProducts()
	defer dbLocalConnection.Close()
	return productList
}
