package main

import (
	"github.com/KKGo-Software-engineering/coaching-session/week-3/dog"
	"github.com/KKGo-Software-engineering/coaching-session/week-3/postgres"
	"github.com/KKGo-Software-engineering/coaching-session/week-3/product"
	"github.com/KKGo-Software-engineering/coaching-session/week-3/speaker"
	"github.com/labstack/echo/v4"
)

func main() {
	speaker.Speak(dog.New())

	p, err := postgres.New()
	if err != nil {
		panic(err)
	}
	defer p.Db.Close()

	e := echo.New()
	handler := product.New(p)
	e.GET("/products/:id", handler.ProductHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
