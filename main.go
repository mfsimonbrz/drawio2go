package main

import (
	"drawio2go/internals"
	"log"

	"github.com/joselitofilho/drawio-parser-go/pkg/parser/xml"
)

type Pet struct {
	Name   string
	Sex    string
	Intact bool
	Age    string
	Breed  string
}

func main() {
	diagram, err := xml.Parse("./erd_3.drawio")
	if err != nil {
		log.Fatal(err)
	}

	elements := internals.BuildElementTree(diagram)
	tables := internals.ProcessElements(elements)

	err = internals.CreateGoModFile("exampleApp", "/Users/marcos.simon/Documents/exampleApp")
	if err != nil {
		log.Fatal(err)
	}

	err = internals.CreateMainFile("exampleApp", tables, "/Users/marcos.simon/Documents/exampleApp")
	if err != nil {
		log.Fatal(err)
	}

	err = internals.CreateModelsFile(tables, "/Users/marcos.simon/Documents/exampleApp/internals/models")
	if err != nil {
		log.Fatal(err)
	}

	err = internals.CreateDBInitFile(tables, "/Users/marcos.simon/Documents/exampleApp/internals/data")
	if err != nil {
		log.Fatal(err)
	}

	err = internals.CreateDataFiles("exampleApp", tables, "/Users/marcos.simon/Documents/exampleApp/internals/data")
	if err != nil {
		log.Fatal(err)
	}

	err = internals.CreateWebFiles("exampleApp", tables, "/Users/marcos.simon/Documents/exampleApp/internals/web")
	if err != nil {
		log.Fatal(err)
	}

}
