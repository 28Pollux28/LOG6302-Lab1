package pretty_print

type Node struct {
	Kind      string
	NChildren int
	Text      string
}

func (n *Node) GetChildrenNumber() int {
	return n.NChildren
}
