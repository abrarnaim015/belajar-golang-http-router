package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T)  {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/files/hello.txt", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, "Hello HttpRouter", string(res))
}