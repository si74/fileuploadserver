package main

import (
	"net/http"
	"fmt"
	"time"
	)

const MAX_UPLOAD_SIZE=1024 *1024

type healthHandler struct {
}

func (h *healthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

type uploadHandler struct {
	infoTime string
}

func (h uploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO(sneha): use object storage data store to hold data vs. saving on disk
	//Info Time
	tm := time.Now().Format(h.infoTime)
	w.Write([]byte("Upload time : \n" + tm))
	//Size of the File for Security
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Support Multiple Files
	files := r.MultipartForm.File["myFile"]
	for _,fileHeader := range files {
		
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}
		//  README  ** Client sends over file name
		fmt.Printf("FileName: %+v\n",fileHeader.Filename)
	}

}

func main() {

	// TODO(sneha): use flags for variables
	addr := ":8080"
	uh :=uploadHandler{ infoTime: time.RFC1123 }
	mux := http.NewServeMux()
	mux.Handle("/health", &healthHandler{})
	mux.Handle("/upload", uh)
	// TODO(sneha): endpoint to retrieve file
	// TODO(sneha): endpoint to delete existing file

	http.ListenAndServe(addr, mux)
}
