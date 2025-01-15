package models

import (
	"bytes"
	"fmt"
	"strings"
)

func (t *Table) AddField(aField *Field) {
	if !t.Contains(aField) {
		t.Fields = append(t.Fields, aField)
	}
}

func (t Table) Contains(aField *Field) bool {
	for _, field := range t.Fields {
		if field.Name == aField.Name {
			return true
		}
	}

	return false
}

func NewTable() *Table {
	return &Table{Name: "", Fields: []*Field{}}
}

func NewField() *Field {
	return &Field{Name: "", Type: "int", Nullable: false, Primary: false, Foreign: false}
}

func (t Table) String() string {
	var result bytes.Buffer
	result.WriteString(t.Name)
	result.WriteString("[")
	for pos, field := range t.Fields {
		if pos > 0 {
			result.WriteString(", " + field.String())
		} else {
			result.WriteString(field.String())
		}
	}
	result.WriteString("]\n")
	return result.String()
}

func (f Field) String() string {
	fieldName := strings.ToLower(strings.Trim(f.Name, ""))
	return fmt.Sprintf("{%s, %s, %t, %t, %t}", fieldName, f.Type, f.Primary, f.Foreign, f.Nullable)
}

func (t Table) GetPrimaryKeyField() *Field {
	for _, field := range t.Fields {
		if field.Primary {
			return field
		}
	}

	return nil
}
