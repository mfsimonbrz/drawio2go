package internals

import (
	"drawio2go/models"
	"errors"
	"fmt"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Element *models.Element
	Nodes   []*Node
}

func NewTree() *Tree {
	aTree := &Tree{}
	aTree.init()

	return aTree
}

func (n *Node) AddChild(newNode *Node) {
	n.Nodes = append(n.Nodes, newNode)
}

func (t *Tree) init() {
	rootElement := Node{Element: nil, Nodes: []*Node{}}
	t.Root = &rootElement
}

func (t Tree) Print() {
	printElements(t.Root)
}

func printElements(n *Node) {
	for _, n := range n.Nodes {
		if len(n.Nodes) > 0 {
			printElements(n)
		} else {
			fmt.Printf("%s\t->\t%s\n", n.Element.Id, n.Element.Value)
		}
	}

	if n.Element != nil {
		fmt.Printf("%s\t->\t%s\n", n.Element.Id, n.Element.Value)
	}
}

func (t Tree) PrintStack() *Stack {
	result := NewStack()
	printStackElements(t.Root, result)

	return result
}

func printStackElements(n *Node, stack *Stack) {
	if n.Element != nil && n.Element.Kind == "table" {
		stopNode := &Node{Element: &models.Element{Value: "stop", Kind: "stop"}}
		stack.Push(stopNode)
	}

	for _, n := range n.Nodes {
		if len(n.Nodes) > 0 {
			printStackElements(n, stack)
		} else {
			stack.Push(n)
		}
	}

	if n.Element != nil {
		stack.Push(n)
	}
}

func (t Tree) getElementById(elementId string, aNode *Node) (*Node, error) {
	if aNode == nil {
		return nil, errors.New("element not found")
	}

	if aNode.Nodes != nil {
		for _, node := range aNode.Nodes {
			found, err := t.getElementById(elementId, node)
			if err == nil {
				return found, nil
			}
		}
	}

	if aNode.Element != nil {
		if aNode.Element.Id == elementId {
			return aNode, nil
		}
	}

	return nil, errors.New("element not found")
}

func (t Tree) GetElementById(elementId string) (*Node, error) {
	return t.getElementById(elementId, t.Root)
}
