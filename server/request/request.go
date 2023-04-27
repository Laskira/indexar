package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Laskira/indexar/server/directories"
	errorsApp "github.com/Laskira/indexar/server/errors"
)

func MappingRequest(w http.ResponseWriter, r *http.Request) {

	dbName := directories.GetDatabaseName()

	url := "http://localhost:4080/api/" + dbName + "/_mapping"
	req, err := http.NewRequest("GET", url, nil)
	errorsApp.ErrorWhatShowAWarning(err)

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	errorsApp.ErrorWhatShowAWarning(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	errorsApp.ErrorWhatShowAWarning(err)

	var jsonData interface{}
	jsonErr := json.Unmarshal([]byte(body), &jsonData)
	if jsonErr != nil {
		panic(jsonErr)
	}

	formattedJson, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	// Write the formatted JSON to the response writer
	w.Write(formattedJson)
}
