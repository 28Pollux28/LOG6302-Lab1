package tree

/*
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
			// if kind is not in the map, set it to 0
			if _, ok := countMap[kind]; !ok {
				countMap[kind] = 0
			}
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

// FindKindTrees finds nodes that match the kind tree map provided
// Returns a map with the kindtree map entry as key and the nodes that match as value
func (n *Node) FindKindTrees(kindTreeMap map[string]KindTree) map[string][]*Node {
	results := n.WalkPostfixWithCallback(n.findKindTrees(kindTreeMap))[0]
	return results.(map[string][]*Node)
}

// A bit more complicated since we have to deal with maps instead of a single slice
// We have to merge the maps from the results (got them from our childs) and then append the nodes that match the kind tree
func (n *Node) findKindTrees(kindTreeMap map[string]KindTree) VisitorFunc {
	return func(n *Node, result []interface{}) []interface{} {
		// Merge Maps from result
		foundNodesMap := make(map[string][]*Node)
		for _, r := range result {
			for k, v := range r.(map[string][]*Node) {
				foundNodesMap[k] = append(foundNodesMap[k], v...)
			}
		}
		// Try to find nodes that match the kind tree attributes
		for kind, kindTree := range kindTreeMap {
			if kindTree.Match(n) {
				foundNodesMap[kind] = append(foundNodesMap[kind], n)
			}
		}
		return []interface{}{foundNodesMap}
	}
}
*/
