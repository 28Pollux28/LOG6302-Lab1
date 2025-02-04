package tree

type VisitorFind struct {
	KindTree KindTree
	Nodes    []*Node
}

type VisitorFinds struct {
	KindTrees map[string]KindTree
	Nodes     map[string][]*Node
}

func (v *VisitorFind) VisitNode(n *Node) {
	if v.KindTree.Match(n) {
		v.Nodes = append(v.Nodes, n)
	}
}

func (v *VisitorFinds) VisitNode(n *Node) {
	for kind, kindtree := range v.KindTrees {
		if kindtree.Match(n) {
			v.Nodes[kind] = append(v.Nodes[kind], n)
		}
	}
}
