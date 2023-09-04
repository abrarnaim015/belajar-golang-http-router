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

func TestRouterPatternNamedParameters(t *testing.T)  {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("id")
		itemId := p.ByName("itemId")
		text := "Product " + productId + " Item " + itemId
		fmt.Fprint(w, text)
	})

	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/products/1/items/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, "Product 1 Item 1", string(res))
}

func TestRouterPatternCatchAllParameters(t *testing.T)  {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		imageCatch := p.ByName("image")
		text := "Image : " + imageCatch
		fmt.Fprint(w, text)
	})

	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/images/small/profile.png", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, "Image : /small/profile.png", string(res))
}