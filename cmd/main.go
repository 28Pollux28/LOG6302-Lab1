package cmd

import (
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/cmd/operations"
	"os"
)

func Main() {
	// This CLI tool will have multiple commands
	// The first command will be to parse a PHP file and output a JSON file with the tree
	// The second command will be to parse a JSON tree file and then perform some operations on it
	// define everything using the flag package

	mainHelp := flag.Bool("help", false, "Show help for the program")

	// Check which command is being run
	flag.Parse()

	if *mainHelp {
		fmt.Println("Usage: ./main [command] [flags]")
		fmt.Println("Commands:")
		fmt.Println("  parse - Parse a PHP file and output a JSON file with the tree")
		fmt.Println("  show - Show the tree of a JSON file")
		fmt.Println("  operations - Input a JSON tree file and then perform some operations on it")
		fmt.Println("Type ./main [command] --help for more information on a command")
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command. Type --help for more information")
		os.Exit(1)
	}
	args := flag.Args()
	switch os.Args[1] {
	case "parse":
		parsePHP(args)
	case "show":
		showTree(args)
	case "operations":
		operations.Main(args)
	default:
		fmt.Println("Please provide a valid command. Type --help for more information")
		os.Exit(1)
	}
	//
	//// Load file name from args
	//if len(os.Args) < 2 {
	//	fmt.Println("Please provide a file name")
	//	os.Exit(1)
	//}
	//fileName := os.Args[1]
	//
	//filePHP, err := os.ReadFile(fileName)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//code := filePHP
	//
	//parser := ts.NewParser()
	//defer parser.Close()
	//parser.SetLanguage(ts.NewLanguage(tree_sitter_php.LanguagePHP()))
	////open, err := os.OpenFile("tree2.dot", os.O_CREATE|os.O_WRONLY, 0644)
	////if err != nil {
	////	fmt.Println(err)
	////	os.Exit(1)
	////}
	////defer open.Close()
	////parser.PrintDotGraphs(open)
	//
	//treesitterTree := parser.Parse(code, nil)
	//defer treesitterTree.Close()
	//
	//root := treesitterTree.RootNode()
	//
	//treeNode := tree.WalkTreeSitterTree(root, &filePHP)
	//treeNode.PrintTree()
	//jsonTree, _ := json.Marshal(treeNode)
	// Write to file
	//file, err := os.Create(fileName + "-tree.json")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//		os.Exit(1)
	//	}
	//}(file)
	//_, err = file.Write(jsonTree)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//fmt.Println(root.ToSexp())
	//countKind := treeNode.CountKind("integer")
	//
	//countKinds := treeNode.CountKinds([]string{"integer", "string_content"})
	//
	//fmt.Printf("Number of integers in the tree: %d\n", countKind)
	//fmt.Printf("Number of integers and string_content in the tree: %v\n", countKinds)
}
