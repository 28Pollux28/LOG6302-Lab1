package tree

import (
	"fmt"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type Point struct {
	Row    uint `json:"row"`
	Column uint `json:"column"`
}

type Attribute[G any] struct {
	V G
}

type Node struct {
	ID             uintptr                   `json:"id"`
	KindId         uint16                    `json:"kind_id"`
	GrammarId      uint16                    `json:"grammar_id"`
	Kind           string                    `json:"kind"`
	GrammarName    string                    `json:"grammar_name"`
	IsNamed        bool                      `json:"is_named"`
	IsExtra        bool                      `json:"is_extra"`
	HasChanges     bool                      `json:"has_changes"`
	HasError       bool                      `json:"has_error"`
	IsError        bool                      `json:"is_error"`
	ParseState     uint16                    `json:"parse_state"`
	NextParseState uint16                    `json:"next_parse_state"`
	IsMissing      bool                      `json:"is_missing"`
	StartByte      uint                      `json:"start_byte"`
	EndByte        uint                      `json:"end_byte"`
	StartPosition  Point                     `json:"start_position"`
	EndPosition    Point                     `json:"end_position"`
	Text           string                    `json:"text"`
	Parent         *Node                     `json:"-"`
	Descendants    []*Node                   `json:"descendants"`
	Node           *ts.Node                  `json:"-"`
	Attributes     map[string]Attribute[any] `json:"attributes"`
}

func NewTreeNode(node *ts.Node, source *[]byte) *Node {
	return &Node{
		ID:             node.Id(),
		KindId:         node.KindId(),
		GrammarId:      node.GrammarId(),
		Kind:           node.Kind(),
		GrammarName:    node.GrammarName(),
		IsNamed:        node.IsNamed(),
		IsExtra:        node.IsExtra(),
		HasChanges:     node.HasChanges(),
		HasError:       node.HasError(),
		IsError:        node.IsError(),
		ParseState:     node.ParseState(),
		NextParseState: node.NextParseState(),
		IsMissing:      node.IsMissing(),
		StartByte:      node.StartByte(),
		EndByte:        node.EndByte(),
		StartPosition:  Point{Row: node.StartPosition().Row, Column: node.StartPosition().Column},
		EndPosition:    Point{Row: node.EndPosition().Row, Column: node.EndPosition().Column},
		Text:           node.Utf8Text(*source),
		Node:           node,
		Parent:         nil,
		Descendants:    []*Node{},
		Attributes:     make(map[string]Attribute[any]),
	}
}

func (n *Node) PrintTree() {
	n.printTree(0)
}

func (n *Node) printTree(level int) {
	if n == nil {
		return
	}

	indent := ""
	for i := 0; i < level; i++ {
		indent += "  |"
	}

	fmt.Printf("%s=>Kind: \"%s\",GrammarName: \"%s\"\n", indent, n.Kind, n.GrammarName)
	for _, child := range n.Descendants {
		child.printTree(level + 1)
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("ID: %d, Kind: %s, GrammarName: %s, IsNamed: %t, IsExtra: %t, HasChanges: %t, HasError: %t, IsError: %t, ParseState: %d, NextParseState: %d, IsMissing: %t, StartByte: %d, EndByte: %d, StartPosition: %v, EndPosition: %v, Text: %s", n.ID, n.Kind, n.GrammarName, n.IsNamed, n.IsExtra, n.HasChanges, n.HasError, n.IsError, n.ParseState, n.NextParseState, n.IsMissing, n.StartByte, n.EndByte, n.StartPosition, n.EndPosition, n.Text)
}
