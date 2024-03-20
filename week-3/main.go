package main

import (
	"database/sql"
	"github.com/KKGo-Software-engineering/coaching-session/week-3/dog"
	"github.com/KKGo-Software-engineering/coaching-session/week-3/postgres"
	"github.com/KKGo-Software-engineering/coaching-session/week-3/product"
	"github.com/KKGo-Software-engineering/coaching-session/week-3/speaker"
	"github.com/labstack/echo/v4"
)

func main() {
	speaker.Speak(dog.New())

	db, err := postgres.New()
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	e := echo.New()
	handler := product.New(db)
	e.GET("/products", handler.ProductHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
