package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/tree"
	ts "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_php "github.com/tree-sitter/tree-sitter-php/bindings/go"

	"os"
)

func parsePHPFile(args []string) {
	parseCmd := flag.NewFlagSet("parse", flag.ExitOnError)
	outputFile := parseCmd.String("output", "tree.json", "The output JSON file")
	prettyPrint := parseCmd.Bool("pretty", false, "Pretty print the JSON output")
	parseHelp := parseCmd.Bool("help", false, "Show help for the parse command")
	parseCmd.Parse(args[1:])

	if *parseHelp {
		fmt.Println("Usage: ./main parse [flags] file.php")
		fmt.Println("Flags:")
		fmt.Println("  --output - The output JSON file")
		fmt.Println("  --pretty - Pretty print the JSON output")
		fmt.Println("  --help - Show help for the parse command")
		os.Exit(0)
	}

	if len(parseCmd.Args()) < 1 {
		fmt.Println("Please provide a file name. Type --help for more information")
		os.Exit(1)
	}
	// Load file name from args
	fileName := parseCmd.Args()[0]
	filePHP, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	code := filePHP

	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(ts.NewLanguage(tree_sitter_php.LanguagePHP()))

	treesitterTree := parser.Parse(code, nil)
	defer treesitterTree.Close()

	root := treesitterTree.RootNode()

	treeNode := tree.WalkTreeSitterTree(root, &filePHP)
	// write to json file
	var jsonTree []byte
	if *prettyPrint {
		jsonTree, err = json.MarshalIndent(treeNode, "", "\t")
	} else {
		jsonTree, err = json.Marshal(treeNode)
	}
	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
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
	os.Exit(0)
}
