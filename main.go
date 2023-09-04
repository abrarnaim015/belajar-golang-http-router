package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main()  {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello Http Router")
	})

	server := http.Server {
		Addr: ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}