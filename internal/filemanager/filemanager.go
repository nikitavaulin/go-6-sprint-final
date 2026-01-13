package filemanager

import (
	"os"
	"time"
)

func GenerateFileName(fileName, extension string) string {
	creationTime := time.Now().Format("02_01_06_15_04_05")
	fileName += "_" + creationTime + extension
	return fileName
}

func CreateFile(path, data string) error {
	return os.WriteFile(path, []byte(data), 0755)
}
