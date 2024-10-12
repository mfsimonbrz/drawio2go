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
	diagram, err := xml.Parse("./erd_1.drawio")
	if err != nil {
		log.Fatal(err)
	}

	elements := internals.BuildElementTree(diagram)
	tables := internals.ProcessElements(elements)

	// // for _, t := range tables {
	// // 	fmt.Println(internals.GenerateStruct(t))
	// // 	fmt.Println()
	// // 	fmt.Println(internals.GenerateCreateTableSQLStatement(t))
	// // 	fmt.Println(internals.GenerateSelectAllStatement(t))
	// // 	fmt.Println(internals.GenerateSelectByIdStatement(t))
	// // 	fmt.Println(internals.GenerateInsertStatement(t))
	// // 	fmt.Println(internals.GenerateUpdateStatement(t))
	// // 	fmt.Println()
	// // 	fmt.Println()
	// // }

	// if _, err := os.Stat("/Users/marcos.simon/Documents/exampleApp"); err != nil {
	// 	os.MkdirAll("/Users/marcos.simon/Documents/exampleApp", 0755)
	// 	os.MkdirAll("/Users/marcos.simon/Documents/exampleApp/internals/models", 0755)
	// 	os.MkdirAll("/Users/marcos.simon/Documents/exampleApp/internals/web", 0755)
	// 	os.MkdirAll("/Users/marcos.simon/Documents/exampleApp/internals/db", 0755)

	// } else {
	// 	fmt.Print("already exists")
	// }

	err = internals.CreateModelsFile(tables, "/Users/marcos.simon/Documents/exampleApp/internals/models")
	if err != nil {
		log.Fatal(err)
	}

}
