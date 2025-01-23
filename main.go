package main

import (
	"encoding/json"
	"fmt"
	"os"

	ts "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_php "github.com/tree-sitter/tree-sitter-php/bindings/go"
)

type Point struct {
	Row    uint `json:"row"`
	Column uint `json:"column"`
}

type TreeNode struct {
	ID             uintptr     `json:"id"`
	KindId         uint16      `json:"kind_id"`
	GrammarId      uint16      `json:"grammar_id"`
	Kind           string      `json:"kind"`
	GrammarName    string      `json:"grammar_name"`
	IsNamed        bool        `json:"is_named"`
	IsExtra        bool        `json:"is_extra"`
	HasChanges     bool        `json:"has_changes"`
	HasError       bool        `json:"has_error"`
	IsError        bool        `json:"is_error"`
	ParseState     uint16      `json:"parse_state"`
	NextParseState uint16      `json:"next_parse_state"`
	IsMissing      bool        `json:"is_missing"`
	StartByte      uint        `json:"start_byte"`
	EndByte        uint        `json:"end_byte"`
	StartPosition  Point       `json:"start_position"`
	EndPosition    Point       `json:"end_position"`
	Text           string      `json:"text"`
	Parent         *TreeNode   `json:"-"`
	Descendants    []*TreeNode `json:"descendants"`
	Node           *ts.Node    `json:"-"`
}

func newTreeNode(node *ts.Node, source *[]byte) *TreeNode {
	return &TreeNode{
		ID:             node.Id(),
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
		Descendants:    []*TreeNode{},
	}
}

func printTree(node *TreeNode, level int) {
	if node == nil {
		return
	}

	indent := ""
	for i := 0; i < level; i++ {
		indent += "  |"
	}

	fmt.Printf("%s=>Kind: %s,GrammarName: %s\n", indent, node.Kind, node.GrammarName)
	for _, child := range node.Descendants {
		printTree(child, level+1)
	}
}

func main() {
	// Load file name from args
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}
	fileName := os.Args[1]

	filePHP, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	code := filePHP

	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(ts.NewLanguage(tree_sitter_php.LanguagePHP()))
	//open, err := os.OpenFile("tree2.dot", os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//defer open.Close()
	//parser.PrintDotGraphs(open)

	tree := parser.Parse(code, nil)
	defer tree.Close()

	root := tree.RootNode()

	treeNode := walkFromNode(root, nil, &filePHP)
	printTree(treeNode, 0)
	jsonTree, _ := json.Marshal(treeNode)
	// Write to file
	file, err := os.Create(fileName + "-tree.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}(file)
	_, err = file.Write(jsonTree)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(root.ToSexp())
}

func walkFromNode(node *ts.Node, parentTreeNode *TreeNode, source *[]byte) *TreeNode {
	if parentTreeNode == nil {
		parentTreeNode = newTreeNode(node, source)
	} else {
		selfTreeNode := newTreeNode(node, source)
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
