package main

import (
	"fmt"
	// "log"
	// "net/http"
	// "net/http/pprof"

	"github.com/Laskira/indexar/server/directories"
	"github.com/Laskira/indexar/server/routes"
	"github.com/Laskira/indexar/server/zincsearch"
)

func main() {

	// mux := http.NewServeMux()
	// mux.HandleFunc("/custom_debug_path/profile", pprof.Profile)
	// log.Fatal(http.ListenAndServe(":7777", mux))

	//Get directorys and converting file in NJson
	directories.GetDirectorysAndConvert()

	zincsearch.SendToZincSearch()

	fmt.Println("Succes files indexing")

	//Chi routing
	routes.StartRouting()

}
