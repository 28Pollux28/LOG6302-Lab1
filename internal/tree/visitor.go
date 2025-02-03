package tree

type VisitorFunc func(n *Node, result []interface{}) []interface{}

func (n *Node) WalkPostfixWithCallback(callback VisitorFunc) []interface{} {
	var results []interface{}
	for _, child := range n.Descendants {
		results = append(results, child.WalkPostfixWithCallback(callback)...)
	}
	return callback(n, results)
}

func (n *Node) WalkPrefixWithCallback(callback VisitorFunc) []interface{} {
	var results []interface{}
	results = append(results, callback(n, results))
	for _, child := range n.Descendants {
		results = append(results, child.WalkPrefixWithCallback(callback)...)
	}
	return results
}
