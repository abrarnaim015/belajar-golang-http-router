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

func TestRouter(t *testing.T)  {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello Http Router")
	})

	req := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, "Hello Http Router", string(res))
}