package main

import "github.com/HoangCaoPhi/go-ecommerce/internals/routers"

func main() {
	r := routers.NewRouter()
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
