package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const uploadDir = "./uploads/"
const port = ":8030"

func main() {
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/files/", getFile)
	http.HandleFunc("/delete/", deleteFile)

	fmt.Printf("Servers started at %s\n", port)
	http.ListenAndServe(port, nil)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[len("/files/"):]
	filePath := filepath.Join(uploadDir, fileName)
	http.ServeFile(w, r, filePath)
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fileName := r.URL.Path[len("/delete/"):]
	filePath := filepath.Join(uploadDir, fileName)

	err := os.Remove(filePath)
	if err != nil {
		http.Error(w, "Error deleting the file", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("File deleted successfully"))
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	r.ParseMultipartForm(10 << 20) // Limit file size to 10 MB

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join(uploadDir, handler.Filename))
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error writing the file", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("File uploaded successfully"))
}
