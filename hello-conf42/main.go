package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		router := spinhttp.NewRouter()
		router.GET("/:kind", handle_get)
		// register more routes

		router.ServeHTTP(w, r)

	})
}

type ResponseModel struct {
	Message string `json:"message"`
}

func handle_get(w http.ResponseWriter, r *http.Request, params spinhttp.Params) {
	k := params.ByName("kind")
	if len(k) == 0 {
		http.Error(w, "Bad Request", 400)
		return
	}
	result := ResponseModel{
		Message: fmt.Sprintf("Hello Conf42 %s!", k),
	}
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err := enc.Encode(result)
	if err != nil {
		http.Error(w, "Error while encoding response message", 500)
		return
	}
}

func main() {}
