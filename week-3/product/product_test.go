package product

import (
	"encoding/json"
	"github.com/KKGo-Software-engineering/coaching-session/week-3/postgres"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProduct(t *testing.T) {
	t.Run("given Refactoring as a product, product name should return Book: Refactoring", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("3")
		db, _ := postgres.New()
		p := New(db.Db)

		p.ProductHandler(c)

		wantProductName := "Book: Refactoring"
		want := DB{Name: wantProductName, Category: "Book", ProductID: 3}
		gotJson := rec.Body.Bytes()
		var got DB
		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal json: %v", err)
		}
		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("given iPhone as a product, product name should return Mobile: iPhone 15 Pro", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		db, _ := postgres.New()
		p := New(db.Db)

		p.ProductHandler(c)

		wantProductName := "Mobile: iPhone 15 Pro"
		want := DB{Name: wantProductName, Category: "Mobile", ProductID: 1}
		gotJson := rec.Body.Bytes()
		var got DB
		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal json: %v", err)
		}
		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}
