package server

import "net/http"

func InterfaceTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
