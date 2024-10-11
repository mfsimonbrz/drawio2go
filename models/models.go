package models

type Element struct {
	Id    string
	Value string
	Kind  string
}

type Field struct {
	Name     string
	Type     string
	Nullable bool
	Primary  bool
	Foreign  bool
}

type Table struct {
	Name    string
	Fields  []*Field
	Imports []string
}
