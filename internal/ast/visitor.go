package ast

type Visitor interface {
	VisitNode(n *Node)
}
