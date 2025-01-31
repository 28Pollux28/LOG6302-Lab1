package operations

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/tree"
	"os"
)

func findKindTrees(fileName string, args []string) {
	// Define the flags for the find-kind-tree operation
	findKindTreesOperation := flag.NewFlagSet("find-kind-tree", flag.ExitOnError)
	findKindTreesHelp := findKindTreesOperation.Bool("help", false, "Show help for the find-kind-tree operation")
	findKindTreesOperation.Parse(args[2:])

	if *findKindTreesHelp {
		// Print help message
		fmt.Println("Usage: ./main operations [flags] file.ast.json find-kind-trees [flags] <kind-trees.kt.json>")
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

	// Load file
	fileJSON, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var treeNode tree.Node
	err = json.Unmarshal(fileJSON, &treeNode)
	if err != nil {
		fmt.Printf("Error parsing tree: %s\n", err)
		os.Exit(1)
	}

	// Load kind tree
	kindTreeJSON, err := os.ReadFile(findKindTreesOperation.Args()[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var kindTrees map[string]tree.KindTree
	err = json.Unmarshal(kindTreeJSON, &kindTrees)
	if err != nil {
		fmt.Printf("Error parsing kind tree: %s\n", err)
		os.Exit(1)
	}

	// Find kind tree in tree
	foundNodesMap := treeNode.FindKindTrees(kindTrees)
	if len(foundNodesMap) == 0 {
		fmt.Println("No nodes found")
		os.Exit(0)
	}
	fmt.Println("Results:")
	for key, nodesArray := range foundNodesMap {
		fmt.Printf("Found occurences for %s : \n", key)
		for _, node := range nodesArray {
			fmt.Printf("file: %s, line: %d\n", fileName, node.StartPosition.Row+1)
		}
		fmt.Print("----------------------\n")
	}

}
