package api

import (
	"encoding/json"
	"fmt"
	"io"
	"lighten.top/lightning/api/service"
	"net/http"
	"time"
)

const MaxFileSize int64 = 2 * 1024 * 1024

func GetIpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	respBody, _ := json.Marshal(map[string]string{
		"ip": service.GetIpFromRequest(r),
	})
	w.Write(respBody)
}

func ShortMessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now())
	b, _ := io.ReadAll(r.Body)
	fmt.Println(string(b))
	w.Write(nil)
}

// UploadFileHandler write a http handler to upload file
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// 1. parse and validate file and post parameters
	err := r.ParseMultipartForm(MaxFileSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// 2. visit files in form
	for _, files := range r.MultipartForm.File {
		for _, file := range files {
			err := service.WriteRequestFileToLocal(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
