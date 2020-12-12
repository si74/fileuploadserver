package main

import "net/http"

type healthHandler struct {
}

func (h *healthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

type uploadHandler struct {
}

func (h *uploadHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO(sneha): use object storage data store to hold data vs. saving on disk
}

func main() {

	// TODO(sneha): use flags for variables
	addr := ":8080"

	mux := http.NewServeMux()
	mux.Handle("/health", &healthHandler{})
	mux.Handle("/upload", &uploadHandler{})

	http.ListenAndServe(addr, mux)
}
