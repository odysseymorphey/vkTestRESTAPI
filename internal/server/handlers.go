package server

import "net/http"

func mock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PENIS"))
}