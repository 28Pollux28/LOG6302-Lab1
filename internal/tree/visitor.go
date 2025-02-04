package tree

type Visitor interface {
	VisitNode(n *Node)
}
