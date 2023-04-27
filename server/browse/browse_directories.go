package browse

import (
	"io/ioutil"

	errorsApp "github.com/Laskira/indexar/server/errors"
	"github.com/Laskira/indexar/server/ndjson"
)

func BrowseDirectories(nameDir string, currentPath string) {

	currentPath += "/" + nameDir

	// get files
	files, err := ioutil.ReadDir(currentPath)
	errorsApp.ErrorWhatStopTheApp(err)

	if len(files) < 1 {
		panic("Files not founds")
	}

	// list with filenames and directories
	var filesList []string
	var directoriesList []string

	for _, file := range files {

		if file.IsDir() {
			directoriesList = append(directoriesList, file.Name())
		} else {
			filesList = append(filesList, file.Name())
		}
	}

	if len(filesList) >= 1 {
		ndjson.ConvertToNDJson(filesList, currentPath)
	}

	for _, dir := range directoriesList {
		BrowseDirectories(dir, currentPath)
	}

	if len(directoriesList) == 0 {
		return
	}
}
