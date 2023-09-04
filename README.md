# belajar-golang-http-router

Link Documentasi <a href="https://github.com/julienschmidt/httprouter">Http Router</a>

Link Pembelajaran <a href="https://www.youtube.com/watch?v=spXgBcjiuWM&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ&index=12&ab_channel=ProgrammerZamanNow">Programmer Zaman Now
</a>

- For http Router

```golang
go get github.com/julienschmidt/httprouter
```

- For Unit Testing

```golang
go get github.com/stretchr/testify
```

How to Get Params

```golang
router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  productId := p.ByName("id") <- how to get params
  text := "Product " + productId
  fmt.Fprint(w, text)
})
```

> Catch All Parameter

- Selain named parameter, ada jg yang bernama `catch all parameter`, yaitu menangkap semua parameter
- `catch all parameter` harus diawali dgn `*(bintang)`, lalu diikuti dengan nama parameter
- `catch a;; parameter` harus berada di posisi akhir `URL`

```text
Pattern <------------> /src/*filepath <------> Params
-> /src/                    ❌                   nil
-> /src/somefile            ✅                 /somefile
-> /stc/subdir/somefile     ✅              /subdir/somefile
```

```golang
router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  productId := p.ByName("id") <- param 1
  itemId := p.ByName("itemId") <- param 2
  text := "Product " + productId + " Item " + itemId
  fmt.Fprint(w, text)
})
```

- Catch All exp

```golang
router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  imageCatch := p.ByName("image")
  text := "Image : " + imageCatch
  fmt.Fprint(w, text)
})

// exp url -> http://localhost:3000/images/small/profile.png
// imageCatch = /small/profile.png
```

> Panic Handler

- Di `Router` sudah disediakan untuk menangani panic, caranga dgn menggunakan attribute
  - `PanicHandler()`
  - ```golang
      func(http.ResponseWriter, *http.Request, interface{})
    ```

```golang
router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
  fmt.Fprint(w, "error : ", error)
}

router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  panic("Ups Panic error")
})

// response -> error : Ups Panic error
```

> Method Not Allowed Handler

- `router.MethodNotAllowed = http.Handler`

```golang
func TestMethodNotAllowed(t *testing.T)  {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Boleh")
	})
	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "POST")
	})

	req := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, "Gak Boleh", string(res))
}
```

> Middleware

```golang
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
```
