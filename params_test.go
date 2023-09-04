package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T)  {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("id")
		text := "Product " + productId
		fmt.Fprint(w, text)
	})

	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/products/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, "Product 1", string(res))
}