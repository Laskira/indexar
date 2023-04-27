package main

import (
	"fmt"

	"github.com/Laskira/indexar/server/directories"
	"github.com/Laskira/indexar/server/routes"
	"github.com/Laskira/indexar/server/zincsearch"
)

func main() {

	//Get directorys and converting file in NJson
	directories.GetDirectorysAndConvert()

	zincsearch.SendToZincSearch()

	fmt.Println("Succes files indexing")

	//Chi routing
	routes.StartRouting()

}
