package tree

func (n *Node) CountKind(kind string) int {
	results := n.WalkPostfixWithCallback(n.countKind(kind))
	return results[0].(int)
}

func (n *Node) CountKinds(kinds []string) map[string]int {
	results := n.WalkPostfixWithCallback(n.countKinds(kinds))
	return results[0].(map[string]int)
}

func (n *Node) countKind(kind string) VisitorFunc {
	return func(n *Node, result []interface{}) []interface{} {
		// Count by increasing an integer
		count := 0
		for _, r := range result {
			count += r.(int)
		}
		if n.Kind == kind {
			count++
		}
		return []interface{}{count}
	}
}

func (n *Node) countKinds(kinds []string) VisitorFunc {
	return func(n *Node, result []interface{}) []interface{} {
		countMap := make(map[string]int)
		for _, r := range result {
			for k, v := range r.(map[string]int) {
				countMap[k] += v
			}
		}
		for _, kind := range kinds {
			if n.Kind == kind {
				countMap[kind]++
			}
		}
		return []interface{}{countMap}
	}
}

func (n *Node) FindKindTree(kindtree KindTree) []*Node {
	results := n.WalkPostfixWithCallback(n.findKindTree(kindtree))
	nodes := make([]*Node, len(results))
	for i, r := range results {
		nodes[i] = r.(*Node)
	}
	return nodes
}

func (n *Node) findKindTree(kindtree KindTree) VisitorFunc {
	return func(n *Node, result []interface{}) []interface{} {
		if result == nil {
			result = []interface{}{}
		}
		// Try to find nodes that match the kind tree attributes
		if kindtree.Match(n) {
			return append(result, n)
		}
		return result
	}
}
