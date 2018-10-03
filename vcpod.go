package main

import (
	"net/http"

	"github.com/himu62/vcpod/api"
)

func main() {
	if err := http.ListenAndServe(":8888", api.NewHandler()); err != nil {
		panic(err)
	}
}
