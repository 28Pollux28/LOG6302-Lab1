package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/tree"
	ts "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_php "github.com/tree-sitter/tree-sitter-php/bindings/go"
	"os"
	"path"
)

func parsePHP(args []string) {
	parseCmd := flag.NewFlagSet("parse", flag.ExitOnError)
	outputFile := parseCmd.String("output", "tree.json", "The output JSON file")
	prettyPrint := parseCmd.Bool("pretty", false, "Pretty print the JSON output")
	directory := parseCmd.Bool("directory", false, "Parse a directory of PHP files")
	recursive := parseCmd.Bool("recursive", false, "Recursively parse a directory of PHP files")
	parseHelp := parseCmd.Bool("help", false, "Show help for the parse command")
	parseCmd.Parse(args[1:])

	if *parseHelp {
		fmt.Println("Usage: ./main parse [flags] <file.php|directory>")
		fmt.Println("Flags:")
		fmt.Println("  --output - The output JSON file / directory")
		fmt.Println("  --pretty - Pretty print the JSON output")
		fmt.Println("  --directory - Parse a directory of PHP files")
		fmt.Println("  --recursive - Recursively parse a directory of PHP files")
		fmt.Println("  --help - Show help for the parse command")
		os.Exit(0)
	}

	if len(parseCmd.Args()) < 1 {
		fmt.Println("Please provide a file name. Type --help for more information")
		os.Exit(1)
	}
	// Load file name from args
	fileName := parseCmd.Args()[0]
	// Check if file is a directory and directory flag is set
	stat, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	if stat.IsDir() && *directory {
		parseDir(fileName, *outputFile, *recursive, *prettyPrint)
		os.Exit(0)
	}

	filePHP, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	parseFile(filePHP, *outputFile, *prettyPrint)
	os.Exit(0)
}

func parseDir(directory, output string, recursive, prettyPrint bool) {
	// Read all files in directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() && recursive {
			parseDir(directory+"/"+file.Name(), output+"/"+file.Name(), recursive, prettyPrint)
		} else if file.IsDir() && !recursive {
			continue
		}
		if path.Ext(file.Name()) != ".php" {
			fmt.Printf("Skipping file %s: not a PHP file\n", file.Name())
			continue
		}
		fmt.Printf("Parsing file %s\n", directory+"/"+file.Name())
		filePHP, err := os.ReadFile(directory + "/" + file.Name())
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file.Name(), err)
			os.Exit(1)
		}
		parseFile(filePHP, output+"/"+file.Name()+".ast.json", prettyPrint)
	}
}

func parseFile(filePHP []byte, outputFile string, prettyPrint bool) {
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(ts.NewLanguage(tree_sitter_php.LanguagePHP()))

	treesitterTree := parser.Parse(filePHP, nil)
	defer treesitterTree.Close()

	root := treesitterTree.RootNode()

	treeNode := tree.WalkTreeSitterTree(root, &filePHP)
	// write to json file
	var jsonTree []byte
	if prettyPrint {
		jsonTree, _ = json.MarshalIndent(treeNode, "", "\t")
	} else {
		jsonTree, _ = json.Marshal(treeNode)
	}
	dir := path.Dir(outputFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0666)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			os.Exit(1)
		}
	}
	file, err := os.Create(outputFile)
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
}
