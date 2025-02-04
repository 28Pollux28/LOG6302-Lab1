package operations

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/28Pollux28/log6302-parser/internal/tree"
	"github.com/28Pollux28/log6302-parser/utils"
)

func findKindTree(fileName string, args []string, directory, recursive bool) {
	// Define the flags for the find-kind-tree operation
	findKindTreeOperation := flag.NewFlagSet("find-kind-tree", flag.ExitOnError)
	findKindTreeHelp := findKindTreeOperation.Bool("help", false, "Show help for the find-kind-tree operation")
	findKindTreeOperation.Parse(args[2:])

	if *findKindTreeHelp {
		// Print help message
		fmt.Println("Usage: go-php-parser operations [OPFlags] <file.ast.json|directory> find-kind-tree [flags] <kind-tree.kt.json>")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the find-kind-tree operation")
		fmt.Println("  <kind-tree.json> - The kind tree to find in the tree")
		fmt.Println("  The kind tree is a JSON file that represents a tree of kinds to find in the tree")
		fmt.Println("  Kind trees are represented as a JSON object with the following structure:")
		fmt.Println("  {")
		fmt.Println("    \"kind\": \"<kind>\",")
		fmt.Println("    \"attributes\": {")
		fmt.Println("      \"<attribute>\": \"<value>\",")
		fmt.Println("      ...")
		fmt.Println("    },")
		fmt.Println("    \"children\": [")
		fmt.Println("      {")
		fmt.Println("        \"kind\": \"<kind>\",")
		fmt.Println("        \"attributes\": {")
		fmt.Println("          \"<attribute>\": \"<value>\",")
		fmt.Println("          ...")
		fmt.Println("        },")
		fmt.Println("        \"children\": [")
		fmt.Println("          ...")
		fmt.Println("        ]")
		fmt.Println("      },")
		fmt.Println("      ...")
		fmt.Println("    ]")
		fmt.Println("  }")
		fmt.Println("  The kind field is the kind of the node to find in the tree")
		fmt.Println("  The attributes field is an object that contains the attributes of the node to find")
		fmt.Println("  Currently, supported attributes are:")
		fmt.Println("    - \"text\": The text value of the node")
		os.Exit(0)
	}

	if len(findKindTreeOperation.Args()) < 1 {
		fmt.Println("Please provide a kind tree. Type --help for more information")
		os.Exit(1)
	}

	// Load kind tree
	kindTreeJSON, err := os.ReadFile(findKindTreeOperation.Args()[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var kindTree tree.KindTree
	err = json.Unmarshal(kindTreeJSON, &kindTree)
	if err != nil {
		fmt.Printf("Error parsing kind tree: %s\n", err)
		os.Exit(1)
	}

	if directory {
		var wg sync.WaitGroup
		findKindTreeDir(fileName, kindTree, recursive, &wg)
		wg.Wait()
		return
	}
	findKindTreeFile(fileName, kindTree)
}

func findKindTreeDir(filename string, kindTree tree.KindTree, recursive bool, wg *sync.WaitGroup) {
	// Read all files in directory
	files, err := os.ReadDir(filename)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() && recursive {
			findKindTreeDir(filename+"/"+file.Name(), kindTree, recursive, wg)
		} else if file.IsDir() && !recursive {
			continue
		}
		if utils.FileExtension(file.Name(), 2) != ".ast.json" {
			continue
		}
		wg.Add(1)
		go func(fileName string, kindTree tree.KindTree) {
			defer wg.Done()
			findKindTreeFile(fileName, kindTree)
		}(filename+"/"+file.Name(), kindTree)
	}
}

func findKindTreeFile(fileName string, kindTree tree.KindTree) {
	// Load file
	fileJSON, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var treeNode tree.Node
	err = json.Unmarshal(fileJSON, &treeNode)
	if err != nil {
		fmt.Printf("Error parsing tree in file %s : %s\n", fileName, err)
		os.Exit(1)
	}

	// Find kind tree in tree
	v := &tree.VisitorFind{KindTree: kindTree}
	treeNode.WalkPostfixWithCallback(v)
	for _, node := range v.Nodes {
		fmt.Printf("%s: found kind tree near line : %d\n", fileName, node.StartPosition.Row+1)
	}
}
