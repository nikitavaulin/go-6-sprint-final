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

const pathIndexHTML = "../index.html"

// GetMainHtmlHandler обрабатывает GET запрос на получение html страницы
func GetMainHtmlHandler(w http.ResponseWriter, request *http.Request) {
	http.ServeFile(w, request, pathIndexHTML)
}

// ConvertFileHandler обрабатывает POST запрос с загруженным файлом для обработки.
//
// Иницирует загрузку результирующего файла на клиенте
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

	outputFilePath, err := saveConvertResult(logger, header, convertedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outputFile, err := os.Open(outputFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("file not found: %s", err), http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	outputFileInfo, err := outputFile.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+outputFile.Name())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprint(outputFileInfo.Size()))
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, outputFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("sending file error: %s", err), http.StatusInternalServerError)
		return
	}
}

// saveConvertResult определяет директорию и путь для сохранения локальных файлов с результатами конвертации
//
// Создает и сохраняет локальный файл. Возвращает сгенерированный путь до файла
func saveConvertResult(logger *log.Logger, fileHeader *multipart.FileHeader, data string) (string, error) {
	directoryPath := filepath.Join("..", "results_archive")
	err := os.MkdirAll(directoryPath, 0755)
	if err != nil && !os.IsExist(err) {
		logger.Println(err.Error())
		return "", err
	}

	ext := filepath.Ext(fileHeader.Filename)
	outputFileName := filemanager.GenerateFileName("convert-result", ext)
	outputFilePath := filepath.Join(directoryPath, outputFileName)

	err = filemanager.CreateFile(outputFilePath, data)
	if err != nil {
		err = fmt.Errorf("output file creation error: %s", err)
		logger.Println(err.Error())
		return "", err
	}
	logger.Printf("Successful: convert result was written in file %s\n", outputFileName)

	return outputFilePath, nil
}
