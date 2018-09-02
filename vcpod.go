package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/himu62/vcpod/api"
	"github.com/himu62/vcpod/secure"
	"github.com/pkg/browser"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", api.NewHandler())
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

	fmt.Printf("%#v\n", mux)

	fmt.Println(`Welcome to vcpod!

Open https://localhost:8888/ to main page
Type Ctrl-C to quit`)

	browser.OpenURL("https://localhost:8888/")
	srv.ListenAndServeTLS("", "")
}
