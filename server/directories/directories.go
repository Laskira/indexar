package directories

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Laskira/indexar/server/browse"
	errorsApp "github.com/Laskira/indexar/server/errors"
	"github.com/Laskira/indexar/server/ndjson"
)

func GetDirectorysAndConvert() {
	dbName := GetDatabaseName()
	currentPath := getCurrentPath() + "/" + dbName
	setDatabaseNameEnvVar(dbName)

	files, err := getFilesInDirectory(currentPath)
	if err != nil {
		panic(err)
	}

	filesList, dirsList := separateFilesAndDirs(files)
	fmt.Println("Files:", filesList)
	fmt.Println("Directories:", dirsList)

	if len(filesList) >= 1 {
		ndjson.ConvertToNDJson(filesList, currentPath)
	}

	for _, dir := range dirsList {
		browse.BrowseDirectories(dir, currentPath)
	}
}

func GetDatabaseName() string {
	args := os.Args
	if len(args) < 2 {
		panic("Database not specified")
	}
	return args[1]
}

func getCurrentPath() string {
	path, err := os.Getwd()
	errorsApp.ErrorWhatStopTheApp(err)
	return path
}

func setDatabaseNameEnvVar(name string) {
	os.Setenv("dbName", name)
}

func getFilesInDirectory(path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if len(files) < 1 {
		return nil, errors.New("files not found")
	}
	return files, nil
}

func separateFilesAndDirs(files []os.FileInfo) ([]string, []string) {
	var filesList []string
	var dirsList []string
	for _, file := range files {
		if file.IsDir() {
			dirsList = append(dirsList, file.Name())
		} else {
			filesList = append(filesList, file.Name())
		}
	}
	return filesList, dirsList
}
