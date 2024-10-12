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

type Module struct {
	Name    string
	Imports []Map
}

type Map struct {
	Key   string
	Value string
}

type Main struct {
	Imports []Map
	Consts  []Map
}

type Models struct {
	HasTimeField bool
	Tables       []*Table
}
