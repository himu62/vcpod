package api

import (
	"net/http"

	"github.com/himu62/vcpod/api/library"
)

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/library/root", library.RootHandler)
	return mux
}
