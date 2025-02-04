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

func findKindTrees(fileName string, args []string, directory, recursive bool) {
	// Define the flags for the find-kind-tree operation
	findKindTreesOperation := flag.NewFlagSet("find-kind-tree", flag.ExitOnError)
	findKindTreesHelp := findKindTreesOperation.Bool("help", false, "Show help for the find-kind-tree operation")
	findKindTreesOperation.Parse(args[2:])

	if *findKindTreesHelp {
		// Print help message
		fmt.Println("Usage: go-php-parser operations [OPFlags] <file.ast.json|directory> find-kind-trees [flags] <kind-trees.kt.json>")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the find-kind-tree operation")
		fmt.Println(" <kind-trees.json> - The kind trees to find in the tree")
		fmt.Println("  The kind trees is a JSON file that represents a map of tree kinds to find in the tree")
		fmt.Println("  Kind trees are represented as a JSON object with the following structure:")
		fmt.Println("  {")
		fmt.Println("    \"map_entry_name\": {")
		fmt.Println("      \"kind\": \"<kind>\",")
		fmt.Println("      \"attributes\": {")
		fmt.Println("        \"<attribute>\": \"<value>\",")
		fmt.Println("        ...")
		fmt.Println("      },")
		fmt.Println("      \"children\": [")
		fmt.Println("        {")
		fmt.Println("          \"kind\": \"<kind>\",")
		fmt.Println("          \"attributes\": {")
		fmt.Println("            \"<attribute>\": \"<value>\",")
		fmt.Println("            ...")
		fmt.Println("          },")
		fmt.Println("          \"children\": [")
		fmt.Println("            ...")
		fmt.Println("          ]")
		fmt.Println("        },")
		fmt.Println("        ...")
		fmt.Println("      ]")
		fmt.Println("    },")
		fmt.Println("    ...")
		fmt.Println("  }")
		fmt.Println("  The kind field is the kind of the node to find in the tree")
		fmt.Println("  The attributes field is an object that contains the attributes of the node to find")
		fmt.Println("  Currently, supported attributes are:")
		fmt.Println("    - \"text\": The text value of the node")
		os.Exit(0)
	}

	if len(findKindTreesOperation.Args()) < 1 {
		fmt.Println("Please provide a kind tree. Type --help for more information")
		os.Exit(1)
	}

	kindTrees := make(map[string]tree.KindTree)
	kindTreesJSON, err := os.ReadFile(findKindTreesOperation.Args()[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(kindTreesJSON, &kindTrees)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if directory {
		var wg sync.WaitGroup
		findKindTreesDir(fileName, kindTrees, recursive, &wg)
		wg.Wait()
		return
	}
	findKindTreesFile(fileName, kindTrees)
}

func findKindTreesDir(directory string, kindTrees map[string]tree.KindTree, recursive bool, wg *sync.WaitGroup) {
	// Read all files in directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() && recursive {
			findKindTreesDir(directory+"/"+file.Name(), kindTrees, recursive, wg)
		} else if file.IsDir() {
			continue
		}
		if utils.FileExtension(file.Name(), 2) != ".ast.json" {
			continue
		}
		wg.Add(1)
		go func(fileName string, kindTrees map[string]tree.KindTree) {
			defer wg.Done()
			findKindTreesFile(fileName, kindTrees)
		}(directory+"/"+file.Name(), kindTrees)
	}
}

func findKindTreesFile(fileName string, kindTrees map[string]tree.KindTree) {
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
	v := &tree.VisitorFinds{
		KindTrees: kindTrees,
		Nodes:     make(map[string][]*tree.Node),
	}
	treeNode.WalkPostfixWithCallback(v)
	if len(v.Nodes) == 0 {
		return
	}
	fmt.Println("Results:")
	for key, nodesArray := range v.Nodes {
		fmt.Printf("Found occurences for %s : \n", key)
		for _, node := range nodesArray {
			fmt.Printf("file: %s, line: %d\n", fileName, node.StartPosition.Row+1)
		}
		fmt.Print("----------------------\n")
	}
}
