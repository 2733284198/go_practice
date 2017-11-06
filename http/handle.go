package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// Handle  实现http操作
type Handle struct {
	dst string
}

func (h *Handle) uploadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if err := h.saveUploadFile(r.Body, h.dst); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func (h *Handle) saveUploadFile(src io.ReadCloser, dst string) error {
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
