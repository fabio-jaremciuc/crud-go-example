package data

import (
	"encoding/json"
	"io"

	_ "github.com/lib/pq"
)

// Product defines the structure of an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type ProductsList struct {
	Items []Product `json:"products"`
}

func (products *ProductsList) AddProduct(product Product) ProductsList {
	products.Items = append(products.Items, product)
	return *products
}

// https://golang.org/pkg/encoding/json/#NewDecoder
func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Product) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *ProductsList) ListToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
