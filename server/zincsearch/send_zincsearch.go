package zincsearch

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	errorsApp "github.com/Laskira/indexar/server/errors"
)

func SendToZincSearch() {

	value, exists := os.LookupEnv("dbName")
	if exists {
		fmt.Println("dbName:", value)
	} else {
		fmt.Println(value, "does not exist, please review the current path of the file")
	}
	fileFound, err := ioutil.ReadFile(os.Getenv("dbName") + ".ndjson")
	errorsApp.ErrorWhatStopTheApp(err)

	h := http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulk", bytes.NewBuffer(fileFound))
	errorsApp.ErrorWhatStopTheApp(err)

	req.SetBasicAuth("admin", "Complexpass#123")
	r, err := h.Do(req)

	errorsApp.ErrorWhatStopTheApp(err)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	errorsApp.ErrorWhatShowAWarning(err)

	fmt.Println(string(body))
}
