package internals

import (
	"drawio2go/models"
	"testing"
)

type FieldTypeTest struct {
	in, expec string
}

var fieldTypeTests = []FieldTypeTest{
	{"nome varchar", "string"},
	{"active boolean", "bool"},
	{"salary double", "float64"},
	{"age integer", "int"},
	{"purchase_date date", "time.Time"},
	{"avocado mango", ""},
}

var databaseFieldTypeTests = []FieldTypeTest{
	{"string", "text"},
	{"float", "double"},
	{"time", "timestamp"},
	{"banana", ""},
}

var fieldNameTests = []FieldTypeTest{
	{"nome varchar", "nome"},
	{"active boolean", "active"},
	{"salary double", "salary"},
	{"age integer", "age"},
	{"purchase_date date", "purchase_date"},
}

func TestGetFieldName(t *testing.T) {
	for _, test := range fieldNameTests {
		if got := GetFieldName(test.in); got != test.expec {
			t.Errorf("Output %q not equal to expected %q", got, test.expec)
		}
	}
}

func TestGetFieldType(t *testing.T) {
	for _, test := range fieldTypeTests {
		if got := GetFieldType(test.in); got != test.expec {
			t.Errorf("Output %q not equal to expected %q", got, test.expec)
		}
	}
}

func TestGetDatabaseFieldType(t *testing.T) {
	for _, test := range databaseFieldTypeTests {
		if got := GetDatabaseFieldType(test.in); got != test.expec {
			t.Errorf("Output %q not equal to expected %q", got, test.expec)
		}
	}
}

func TestHasTimeField(t *testing.T) {
	tableOne := models.NewTable()
	fieldName := models.NewField()
	fieldName.Name = "name"
	fieldName.Type = "string"
	fieldAge := models.NewField()
	fieldAge.Name = "age"
	fieldAge.Type = "time.Time"
	tableOne.AddField(fieldName)
	tableTwo := models.NewTable()
	tableTwo.AddField(fieldAge)
	tables := []*models.Table{tableOne, tableTwo}

	got := HasTimeField(tables)
	expected := true

	if got != expected {
		t.Errorf("Output %t not equal to expected %t", got, expected)
	}

}
