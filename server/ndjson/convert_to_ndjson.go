package ndjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	errorsApp "github.com/Laskira/indexar/server/errors"
)

func ConvertToNDJson(fileNames []string, path string) {
	// Extract the file name index from the path
	splitPath := strings.Split(path, "/")
	var nameIndex string
	if len(splitPath) >= 2 {
		nameIndex = strings.TrimPrefix(splitPath[len(splitPath)-2], "_") + "." + splitPath[len(splitPath)-1]
	} else {
		nameIndex = strings.TrimPrefix(splitPath[len(splitPath)-1], "_")
	}

	// Calculate total size of files
	var totalSize int64
	for _, fileName := range fileNames {
		if fileInfo, err := os.Stat(path + "/" + fileName); err == nil {
			totalSize += fileInfo.Size()
		} else {
			fmt.Printf("File does not exist: %s\n", fileName)
		}
	}

	// Split fileNames into smaller chunks if total size is too big
	if totalSize > 700000 {
		chunkSlice(fileNames, len(fileNames)/2, path)
		return
	}

	// Create first dictionary for NDJSON format
	firstDict := map[string]map[string]string{
		"index": {
			"_index": os.Getenv("dbName"),
		},
	}
	firstJSON, err := json.Marshal(firstDict)
	errorsApp.ErrorWhatStopTheApp(err)

	// Create second dictionary for NDJSON format
	secondDict := make(map[string]string)
	for _, fileName := range fileNames {
		fileContent, err := ioutil.ReadFile(path + "/" + fileName)
		errorsApp.ErrorWhatStopTheApp(err)
		secondDict[nameIndex+"."+fileName] = string(fileContent)
	}
	secondJSON, err := json.Marshal(secondDict)
	errorsApp.ErrorWhatStopTheApp(err)

	// Write the two JSON dictionaries to NDJSON file
	writeFile(firstJSON, secondJSON)
}

// Split slice into smaller chunks and call ConvertToNDJson on each chunk
func chunkSlice(slice []string, chunkSize int, path string) {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	for _, chunk := range chunks {
		ConvertToNDJson(chunk, path)
	}
}

// Write the two JSON dictionaries to an NDJSON file
func writeFile(firstJSON []byte, secondJSON []byte) {
	fileName := os.Getenv("dbName") + ".ndjson"
	var file *os.File
	var err error
	if _, err = os.Stat(fileName); err == nil {
		// If the file exists, append to it
		file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0660)
		errorsApp.ErrorWhatStopTheApp(err)
	} else {
		// If the file does not exist, create it
		file, err = os.Create(fileName)
		errorsApp.ErrorWhatStopTheApp(err)
	}
	defer file.Close()

	// Write the two JSON dictionaries to the file
	fmt.Fprintln(file, string(firstJSON))
	fmt.Fprintln(file, string(secondJSON))
}
