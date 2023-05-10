package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Laskira/indexar/server/directories"
	errorsApp "github.com/Laskira/indexar/server/errors"
	"github.com/Laskira/indexar/server/models"
)

func DataRequest(w http.ResponseWriter, r *http.Request) {
	var search models.Search
	dbName := directories.GetDatabaseName()

	keyCharacters := search.Searching
	query := ""

	if len(keyCharacters) == 0 {
		query = `{"query": {"match_all": {}}, "size":10}`
	} else {
		query = fmt.Sprintf(`{"query": {"match": {"_all": "%s" }}, "size":10}`, keyCharacters)
	}

	req, err := http.NewRequest("POST", "http://localhost:4080/es/"+dbName+"/_search", strings.NewReader(query))
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
