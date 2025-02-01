package operations

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/tree"
	"os"
	"sync"
)

func countKinds(fileName string, args []string) {
	countKindsOperation := flag.NewFlagSet("count-kinds", flag.ExitOnError)
	countKindsHelp := countKindsOperation.Bool("help", false, "Show help for the count-kind operation")
	countKindsRecursive := countKindsOperation.Bool("recursive", false, "Recursively count the kinds in a directory of AST trees")
	countKindsOperation.Parse(args[2:])

	if *countKindsHelp {
		fmt.Println("Usage: ./main operations [flags] file.ast.json count-kinds [flags] <kind1> <kind2> ...")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the count-kinds operation")
		fmt.Println("  --recursive - Recursively count the kinds in a directory of AST trees")
		fmt.Println(" <kind1> <kind2> ... - The kinds of nodes to count. Refer to the PHP tree-sitter grammar for the kinds")
		os.Exit(0)
	}

	if len(countKindsOperation.Args()) < 1 {
		fmt.Println("Please provide at least one kind. Type --help for more information")
		os.Exit(1)
	}

	if *countKindsRecursive {
		var wg sync.WaitGroup
		countKindsDir(fileName, countKindsOperation.Args(), &wg)
		wg.Wait()
		return
	}
	countKindsFile(fileName, countKindsOperation.Args())
}

func countKindsDir(directory string, kinds []string, wg *sync.WaitGroup) {
	// Read all files in directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() {
			countKindsDir(directory+"/"+file.Name(), kinds, wg)
		} else if !file.IsDir() {
			wg.Add(1)
			go func(fileName string, kind []string) {
				defer wg.Done()
				countKindsFile(fileName, kinds)
			}(directory+"/"+file.Name(), kinds)
		}
	}
}

func countKindsFile(fileName string, kinds []string) {
	// Load file
	fileJSON, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var treeNode tree.Node
	err = json.Unmarshal(fileJSON, &treeNode)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	counts := treeNode.CountKinds(kinds)
	fmt.Printf("Results for file %s\n", fileName)
	for kind, count := range counts {
		fmt.Printf("%s: %d\n", kind, count)
	}
}
