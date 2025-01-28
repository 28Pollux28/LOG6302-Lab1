package tree

import ts "github.com/tree-sitter/go-tree-sitter"

func WalkTreeSitterTree(node *ts.Node, source *[]byte) *Node {
	return walkFromNode(node, nil, source)
}

func walkFromNode(node *ts.Node, parentTreeNode *Node, source *[]byte) *Node {
	if parentTreeNode == nil {
		parentTreeNode = NewTreeNode(node, source)
	} else {
		selfTreeNode := NewTreeNode(node, source)
		parentTreeNode.Descendants = append(parentTreeNode.Descendants, selfTreeNode)
		selfTreeNode.Parent = parentTreeNode
		parentTreeNode = selfTreeNode
	}
	walk := node.Walk()
	defer walk.Close()

	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(uint(i))
		walkFromNode(child, parentTreeNode, source)
	}
	return parentTreeNode
}
