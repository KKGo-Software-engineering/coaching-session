package product

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Product struct {
	store Storer
}

type DB struct {
	ProductID int    `postgres:"product_id" json:"product_id"`
	Name      string `postgres:"name" json:"name"`
	Category  string `postgres:"category" json:"category"`
}

func New(store Storer) Product {
	return Product{store: store}
}

type Storer interface {
	ProductById(id string) (DB, error)
}

type Err struct {
	Message string `json:"message"`
}

func (p Product) ProductHandler(c echo.Context) error {
	id := c.Param("id")
	product, err := p.store.ProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	if product.ProductID == 0 {
		return c.JSON(http.StatusNotFound, Err{Message: "product not found"})
	}
	if product.Category == "Book" {
		product.Name = "Book: " + product.Name
	}
	if product.Category == "Mobile" {
		product.Name = "Mobile: " + product.Name
	}
	return c.JSON(http.StatusOK, product)
}
