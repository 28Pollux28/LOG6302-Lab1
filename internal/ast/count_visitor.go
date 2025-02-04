package ast

type VisitorCount struct {
	Count int
	Kind  string
}

type VisitorCounts struct {
	Counts map[string]int
	Kinds  []string
}

func (v *VisitorCount) VisitNode(n *Node) {
	if n.Kind == v.Kind {
		v.Count++
	}
}

func (v *VisitorCounts) VisitNode(n *Node) {
	for _, kind := range v.Kinds {
		if n.Kind == kind {
			v.Counts[kind]++
		}
	}
}
