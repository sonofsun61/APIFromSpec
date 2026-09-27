package handler

import "net/http"

func setCacheControl(w http.ResponseWriter, value string) {
	w.Header().Set("Cache-Control", value)
}