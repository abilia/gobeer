package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const uploadPath = "pictures/"

// DOES NOT WORK ON SERVER (prematurely closed connection)
func uploadFile(respWriter http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, _, err := request.FormFile("file")
	if err != nil {
		respWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileName, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	id := fileName.String()
	storePath := filepath.Join(uploadPath, id)
	newFile, err := os.Create(storePath)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	newFile.Write(fileBytes)

	var picture = Picture{UUID: id}
	respWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respWriter).Encode(picture)
}
