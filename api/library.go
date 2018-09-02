package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/himu62/vcpod/object"
)

func libraryRootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listRootObjects(w)
	}
}

func listRootObjects(w http.ResponseWriter) {
	roots := make([]object.RootObject, 1)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(roots); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
}
