package product

import (
	"github.com/KKGo-Software-engineering/coaching-session/week-3/postgres"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Product struct {
	*postgres.Postgres
}

type DB struct {
	ProductID int    `postgres:"product_id"`
	Name      string `postgres:"name"`
}

func New(db *postgres.Postgres) Product {
	return Product{db}
}

type Storer interface {
	Products() ([]DB, error)
}

func (p Product) ProductHandler(c echo.Context) error {
	rows, err := p.Db.Query("SELECT product_id,name FROM product")
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
