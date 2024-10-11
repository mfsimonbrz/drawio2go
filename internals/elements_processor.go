package internals

import (
	"drawio2go/models"
	"strings"

	"github.com/joselitofilho/drawio-parser-go/pkg/parser/xml"
)

func BuildElementTree(diagram *xml.MxFile) *Tree {
	elements := NewTree()

	for _, mxCell := range diagram.Diagram.MxGraphModel.Root.MxCells {
		if mxCell.Value != "" && strings.Contains(mxCell.Style, "shape=table") {
			table := &models.Element{Id: mxCell.ID, Value: mxCell.Value, Kind: "table"}
			node := &Node{Element: table}
			elements.Root.AddChild(node)
		}
	}

	for _, mxCell := range diagram.Diagram.MxGraphModel.Root.MxCells {
		if mxCell.Parent != "" {
			node, _ := elements.GetElementById(mxCell.Parent)
			if node != nil {
				child := &models.Element{Id: mxCell.ID, Value: mxCell.Value, Kind: "field"}
				newNode := &Node{Element: child}
				node.AddChild(newNode)
			}
		}
	}

	return elements
}

func ProcessElements(elements *Tree) []*models.Table {
	stack := elements.PrintStack()
	sn := stack.Pop()
	var table *models.Table
	field := models.NewField()
	var tables []*models.Table
	for sn != nil {
		elem := sn.Node.Element
		if elem != nil && elem.Value != "" {
			if elem.Kind == "table" {
				table = models.NewTable()
				table.Name = strings.ToLower(elem.Value)
			} else {
				if elem.Kind != "stop" {
					if !(strings.Contains(elem.Value, "PK") || strings.Contains(elem.Value, "FK")) {
						if field.Name != "" {
							field = models.NewField()
						}
						field.Name = GetFieldName(elem.Value)
						field.Type = GetFieldType(elem.Value)
						table.AddField(field)
						if field.Type == "time.Time" {
							table.Imports = append(table.Imports, "time.Time")
						}

						if strings.Contains(elem.Value, "NOT NULL") {
							field.Nullable = false
						}
					} else {
						if strings.Contains(elem.Value, "PK") {
							field.Primary = true
						} else if strings.Contains(elem.Value, "FK") {
							field.Foreign = true
						}
					}
				} else {
					tables = append(tables, table)
				}
			}
		}
		sn = stack.Pop()
	}

	return tables
}
