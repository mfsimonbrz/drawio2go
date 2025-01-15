package main

import (
	"drawio2go/internals"
	"fmt"
	"log"

	"github.com/joselitofilho/drawio-parser-go/pkg/parser/xml"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{}
	root.AddCommand(Hello())

	root.Execute()
}

func Hello() *cobra.Command {
	return &cobra.Command{
		Use:   "generate [app-name] [input] [output]",
		Short: "Generates the source files based on the input ERD and saves in the output folderß",
		Args:  cobra.ExactArgs(3),
		Run:   generateSources,
	}
}

func generateSources(cmd *cobra.Command, args []string) {
	appName := args[0]
	erdFile := args[1]
	outputDir := args[2]
	output := fmt.Sprintf("%s/%s", outputDir, appName)

	diagram, err := xml.Parse(erdFile)
	if err != nil {
		log.Fatal(err)
	}

	elements := internals.BuildElementTree(diagram)
	tables := internals.ProcessElements(elements)

	err = internals.CreateGoModFile(appName, output)
	if err != nil {
		log.Fatal(err)
	}

	err = internals.CreateMainFile(appName, tables, output)
	if err != nil {
		log.Fatal(err)
	}

	err = internals.CreateModelsFile(tables, fmt.Sprintf("%s/internals/models", output))
	if err != nil {
		log.Fatal(err)
	}

	err = internals.CreateDBInitFile(tables, fmt.Sprintf("%s/internals/data", output))
	if err != nil {
		log.Fatal(err)
	}

	err = internals.CreateDataFiles(appName, tables, fmt.Sprintf("%s/internals/data", output))
	if err != nil {
		log.Fatal(err)
	}

	err = internals.CreateWebFiles(appName, tables, fmt.Sprintf("%s/internals/web", output))
	if err != nil {
		log.Fatal(err)
	}
}
