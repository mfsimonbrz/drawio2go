package internals

type Stack struct {
	top *StackNode
}

func NewStack() *Stack {
	return &Stack{top: nil}
}

type StackNode struct {
	Node *Node
	Next *StackNode
}

func (s *Stack) Push(newNode *Node) {
	newStackNode := &StackNode{Node: newNode, Next: nil}
	if s.top == nil {
		s.top = newStackNode
	} else {
		newStackNode.Next = s.top
		s.top = newStackNode
	}
}

func (s *Stack) Pop() *StackNode {
	result := s.top
	if s.top != nil {
		s.top = s.top.Next
	}

	return result
}

func (s *Stack) Peek() *StackNode {
	return s.top
}
