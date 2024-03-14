package server

import "net/http"

func mock(w http.ResponseWriter, r *http.Request) {
	r := w.
	w.Write([]byte("PENIS"))
}