package internals

import (
	"drawio2go/models"
	"testing"
)

func TestGetElementById(t *testing.T) {
	tree := NewTree()
	elem_1 := &models.Element{Id: "1", Value: "1"}
	node_1 := &Node{Element: elem_1}

	elem_2 := &models.Element{Id: "2", Value: "2"}
	node_2 := &Node{Element: elem_2}

	elem_3 := &models.Element{Id: "3", Value: "3"}
	node_3 := &Node{Element: elem_3}

	elem_4 := &models.Element{Id: "4", Value: "4"}
	node_4 := &Node{Element: elem_4}

	elem_5 := &models.Element{Id: "5", Value: "5"}
	node_5 := &Node{Element: elem_5}

	elem_6 := &models.Element{Id: "6", Value: "6"}
	node_6 := &Node{Element: elem_6}

	tree.Root.AddChild(node_1)
	tree.Root.AddChild(node_2)
	tree.Root.AddChild(node_3)
	node_2.AddChild(node_5)
	tree.Root.AddChild(node_4)
	node_5.AddChild(node_6)

	got, _ := tree.GetElementById("4")
	expected := node_4

	if got != expected {
		t.Errorf("Output %q not equal to expected %q", got.Element.Id, expected.Element.Id)
	}
}

func TestPrintStack(t *testing.T) {
	tree := NewTree()
	elem_1 := &models.Element{Id: "1", Value: "1"}
	node_1 := &Node{Element: elem_1}

	elem_2 := &models.Element{Id: "2", Value: "2"}
	node_2 := &Node{Element: elem_2}

	elem_3 := &models.Element{Id: "3", Value: "3"}
	node_3 := &Node{Element: elem_3}

	elem_4 := &models.Element{Id: "4", Value: "4"}
	node_4 := &Node{Element: elem_4}

	elem_5 := &models.Element{Id: "5", Value: "5"}
	node_5 := &Node{Element: elem_5}

	elem_6 := &models.Element{Id: "6", Value: "6"}
	node_6 := &Node{Element: elem_6}

	tree.Root.AddChild(node_1)
	tree.Root.AddChild(node_2)
	tree.Root.AddChild(node_3)
	node_2.AddChild(node_5)
	tree.Root.AddChild(node_4)
	node_5.AddChild(node_6)

	stack := tree.PrintStack()

	testCases := []string{"4", "3", "2", "5"}

	for ix := 0; ix < len(testCases); ix++ {
		got := stack.Pop()
		if got.Node.Element.Value != testCases[ix] {
			t.Errorf("Output %q not equal to expected %q", got.Node.Element.Value, testCases[ix])
		}
	}
}
