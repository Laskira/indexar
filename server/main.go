package main

import (
	"fmt"

	"github.com/Laskira/indexar/server/routes"
)

func main() {

	//Get directorys and converting file in NJson
	// directories.GetDirectorysAndConvert()

	// zincsearch.SendToZincSearch()

	fmt.Println("Succes files indexing")

	//Chi routing
	routes.StartRouting()

}
