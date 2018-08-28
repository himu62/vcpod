package main

import (
	"fmt"
	"net/http"
	"github.com/pkg/browser"
)

type apiHandler struct{}
func (h *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "api mock")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", &apiHandler{})
	mux.Handle("/", http.FileServer(http.Dir("public_html")))

	fmt.Println(`Welcome to vcpod!

Open http://localhost:8888/ to main page
Type Ctrl-C to quit`)

	browser.OpenURL("http://localhost:8888/")
	http.ListenAndServe(":8888", mux)
}
