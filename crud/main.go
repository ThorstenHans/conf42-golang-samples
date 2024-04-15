package main

import (
	"fmt"
	"net/http"

	"github.com/ThorstenHans/conf42-golang-samples/crud/pkg/api"
	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	fmt.Println("Hello")
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		router := api.New()
		router.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Test")
}
