package product

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubProduct struct {
	product DB
	err     error
}

func (s StubProduct) ProductById(id string) (DB, error) {
	return s.product, s.err
}

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

		stubRefactoring := StubProduct{
			product: DB{Name: "Refactoring", Category: "Book", ProductID: 3},
		}
		p := New(stubRefactoring)

		p.ProductHandler(c)

		wantProductName := "Book: Refactoring"
		want := DB{Name: wantProductName, Category: "Book", ProductID: 3}
		gotJson := rec.Body.Bytes()
		var got DB
		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal json: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
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

		stubIphone := StubProduct{
			product: DB{Name: "iPhone 15 Pro", Category: "Mobile", ProductID: 1},
		}
		p := New(stubIphone)

		p.ProductHandler(c)

		wantProductName := "Mobile: iPhone 15 Pro"
		want := DB{Name: wantProductName, Category: "Mobile", ProductID: 1}
		gotJson := rec.Body.Bytes()
		var got DB
		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal json: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("given unable to connect to database should return Internal Server Error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("10")

		stubError := StubProduct{err: echo.ErrInternalServerError}
		p := New(stubError)

		p.ProductHandler(c)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})
}
