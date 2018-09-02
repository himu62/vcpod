package api

import (
	"net/http"
)

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/library/root", libraryRootHandler)
	return mux
}
