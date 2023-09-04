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

func TestPanicHandler(t *testing.T)  {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(w, "error : ", error)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups Panic error")
	})

	req := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, "error : Ups Panic error", string(res))
}