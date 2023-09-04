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

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Receive Request")
	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T)  {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	middleware := LogMiddleware {
		Handler: router,
	}

	req := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	rec := httptest.NewRecorder()

	middleware.ServeHTTP(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, "Middleware", string(res))
}