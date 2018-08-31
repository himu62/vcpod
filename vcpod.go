package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/himu62/vcpod/secure"
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

	cert, err := secure.CreateDummyCert()
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:      ":8888",
		Handler:   mux,
		TLSConfig: &tls.Config{Certificates: cert},
	}

	fmt.Println(`Welcome to vcpod!

Open https://localhost:8888/ to main page
Type Ctrl-C to quit`)

	browser.OpenURL("https://localhost:8888/")
	srv.ListenAndServeTLS("", "")
}
