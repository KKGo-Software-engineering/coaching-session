package product

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Product struct {
	store Storer
	db    *sql.DB
}

type DB struct {
	ProductID int    `postgres:"product_id" json:"product_id"`
	Name      string `postgres:"name" json:"name"`
	Category  string `postgres:"category" json:"category"`
}

func New(db *sql.DB) Product {
	return Product{db: db}
}

type Storer interface {
	Products() (DB, error)
}

type Err struct {
	Message string `json:"message"`
}

func (p Product) ProductHandler(c echo.Context) error {
	id := c.Param("id")
	rows, err := p.db.Query("SELECT product_id,name,category FROM product WHERE product_id = $1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	defer rows.Close()
	var product DB
	for rows.Next() {
		var p DB
		err = rows.Scan(&p.ProductID, &p.Name, &p.Category)
		if err != nil {
			log.Fatal(err)
			return err
		}
		product = p
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
