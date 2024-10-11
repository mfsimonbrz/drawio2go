package internals

import (
	"drawio2go/models"
	"testing"
)

func TestPush(t *testing.T) {
	element_1 := &models.Element{Id: "1", Value: "", Kind: ""}
	node_1 := &Node{Element: element_1}

	element_2 := &models.Element{Id: "2", Value: "", Kind: ""}
	node_2 := &Node{Element: element_2}

	element_3 := &models.Element{Id: "3", Value: "", Kind: ""}
	node_3 := &Node{Element: element_3}

	stack := NewStack()
	stack.Push(node_1)
	stack.Push(node_2)
	stack.Push(node_3)

	got := stack.Peek().Node.Element.Id
	expec := "3"

	if got != expec {
		t.Errorf("Output %q not equal to expected %q", got, expec)
	}
}

func TestPop(t *testing.T) {
	element_1 := &models.Element{Id: "1", Value: "", Kind: ""}
	node_1 := &Node{Element: element_1}

	element_2 := &models.Element{Id: "2", Value: "", Kind: ""}
	node_2 := &Node{Element: element_2}

	element_3 := &models.Element{Id: "3", Value: "", Kind: ""}
	node_3 := &Node{Element: element_3}

	element_4 := &models.Element{Id: "4", Value: "", Kind: ""}
	node_4 := &Node{Element: element_4}

	stack := NewStack()
	stack.Push(node_1)
	stack.Push(node_2)
	stack.Push(node_3)
	stack.Push(node_4)

	got := stack.Pop().Node.Element.Id
	expec := "4"

	if got != expec {
		t.Errorf("Output %q not equal to expected %q", got, expec)
	}
}
