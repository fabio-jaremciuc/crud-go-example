package handlers

import (
	"log"
	"net/http"

	"github.com/product-api/data"
	"github.com/product-api/repository"
)

type LoggingProducts struct {
	l *log.Logger
}

func NewLoggingProducts(l *log.Logger) *LoggingProducts {
	return &LoggingProducts{l}
}

func (p *LoggingProducts) GetLoggingProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET LoggingProducts")

	// serialize the list to JSON
	productList := data.ProductsList{}
	productList = repository.GetAllProducts()

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	productList.ListToJSON(rw)

}

func (p *LoggingProducts) PostLoggingProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST LoggingProducts")

	// serialize the list to JSON
	productList := data.ProductsList{}
	productList = repository.InsertNewProduct(r)

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	productList.ListToJSON(rw)

}

func (p *LoggingProducts) DeleteLoggingProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE LoggingProducts")

	// serialize the list to JSON
	productList := data.ProductsList{}
	productList = repository.DeleteProduct(r)

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	productList.ListToJSON(rw)

}

func (p *LoggingProducts) UpdateLoggingProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle UPDATE LoggingProducts")

	// serialize the list to JSON
	productList := data.ProductsList{}
	productList = repository.DeleteProduct(r)

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	productList.ListToJSON(rw)

}

func (p *LoggingProducts) PopulateDataLogging(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Populating the database!")

	// serialize the list to JSON
	productList := data.ProductsList{}
	productList = repository.InitialData()

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	productList.ListToJSON(rw)

}
