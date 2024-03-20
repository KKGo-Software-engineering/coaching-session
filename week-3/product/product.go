package product

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Product struct {
	db *sql.DB
}

type DB struct {
	ProductID int    `postgres:"product_id"`
	Name      string `postgres:"name"`
}

func New(db *sql.DB) Product {
	return Product{db: db}
}

func (p Product) ProductHandler(c echo.Context) error {
	rows, err := p.db.Query("SELECT product_id,name FROM product")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	var products []DB
	for rows.Next() {
		var p DB
		err = rows.Scan(&p.ProductID, &p.Name)
		if err != nil {
			log.Fatal(err)
			return err
		}
		products = append(products, p)
	}
	return c.JSON(http.StatusOK, products)
}
