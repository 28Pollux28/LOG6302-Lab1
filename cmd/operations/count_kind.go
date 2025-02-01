package operations

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/tree"
	"os"
	"sync"
)

func countKind(fileName string, args []string) {
	countKindOperation := flag.NewFlagSet("count-kind", flag.ExitOnError)
	countKindHelp := countKindOperation.Bool("help", false, "Show help for the count-kind operation")
	countKindRecursive := countKindOperation.Bool("recursive", false, "Recursively count the kind in a directory of AST trees")
	countKindOperation.Parse(args[2:])

	if *countKindHelp {
		fmt.Println("Usage: ./main operations [flags] file.ast.json count-kind [flags] <kind>")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the count-kind operation")
		fmt.Println("  --recursive - Recursively count the kind in a directory of AST trees")
		fmt.Println("  <kind> - The kind of node to count. Refer to the PHP tree-sitter grammar for the kinds")
		os.Exit(0)
	}

	if len(countKindOperation.Args()) < 1 {
		fmt.Println("Please provide a kind. Type --help for more information")
		os.Exit(1)
	}

	if *countKindRecursive {
		var wg sync.WaitGroup
		countKindDir(fileName, countKindOperation.Args()[0], &wg)
		wg.Wait()
		return
	}
	countKindFile(fileName, countKindOperation.Args()[0])
	return
}

func countKindDir(directory, kind string, wg *sync.WaitGroup) {
	// Read all files in directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() {
			countKindDir(directory+"/"+file.Name(), kind, wg)
		} else if !file.IsDir() {
			wg.Add(1)
			go func(fileName, kind string) {
				defer wg.Done()
				countKindFile(fileName, kind)
			}(directory+"/"+file.Name(), kind)
		}
	}
}

func countKindFile(fileName, kind string) {
	// Load file
	fileJSON, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var treeNode tree.Node
	err = json.Unmarshal(fileJSON, &treeNode)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON in file %s: %v\n", fileName, err)
		os.Exit(1)
	}
	count := treeNode.CountKind(kind)
	fmt.Printf("%s : Number of nodes of kind %s: %d\n", fileName, kind, count)
}
