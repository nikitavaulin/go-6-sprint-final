package handlers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/filemanager"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// const pathIndexHTML = "github.com/Yandex-Practicum/go1fl-sprint6-final/index.html"
const pathIndexHTML = "D:/projects/go-6-sprint-final/index.html"

func GetMainHtmlHandler(w http.ResponseWriter, request *http.Request) {
	http.ServeFile(w, request, pathIndexHTML)
}

func ConvertFileHandler(logger *log.Logger, w http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, fmt.Sprintf("parse form error: %s", err), http.StatusInternalServerError)
		return
	}

	file, header, err := request.FormFile("myFile")
	if err != nil {
		http.Error(w, fmt.Sprintf("getting file error: %s", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("read file error: %s", err), http.StatusInternalServerError)
		return
	}
	inputData := string(buf)

	convertedData := service.ConvertText(inputData)
	saveConvertResult(logger, header, convertedData)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertedData))
}

func saveConvertResult(logger *log.Logger, fileHeader *multipart.FileHeader, data string) {
	directoryPath := filepath.Join("..", "results_archive")
	err := os.MkdirAll(directoryPath, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
		return
	}

	ext := filepath.Ext(fileHeader.Filename)
	outputFileName := filemanager.GenerateFileName("convert-result", ext)
	outputFilePath := filepath.Join(directoryPath, outputFileName)

	err = filemanager.CreateFile(outputFilePath, data)
	if err != nil {
		logger.Printf("output file creation error: %s", err)
		return
	}
	logger.Printf("Successful: convert result was written in file %s", outputFileName)
}
