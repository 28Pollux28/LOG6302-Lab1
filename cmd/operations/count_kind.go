package operations

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/28Pollux28/log6302-parser/internal/ast"
	"github.com/28Pollux28/log6302-parser/utils"
)

func countKind(fileName string, args []string, directory, recursive bool) {
	countKindOperation := flag.NewFlagSet("count-kind", flag.ExitOnError)
	countKindHelp := countKindOperation.Bool("help", false, "Show help for the count-kind operation")
	countKindOperation.Parse(args[2:])

	if *countKindHelp {
		fmt.Println("Usage: go-php-parser operations <file.ast.json|directory> count-kind [flags] <kind>")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the count-kind operation")
		fmt.Println("  <kind> - The kind of node to count. Refer to the PHP tree-sitter grammar for the kinds")
		os.Exit(0)
	}

	if len(countKindOperation.Args()) < 1 {
		fmt.Println("Please provide a kind. Type --help for more information")
		os.Exit(1)
	}

	if directory {
		var wg sync.WaitGroup
		countKindDir(fileName, countKindOperation.Args()[0], recursive, &wg)
		wg.Wait()
		return
	}
	countKindFile(fileName, countKindOperation.Args()[0])
	return
}

func countKindDir(directory, kind string, recursive bool, wg *sync.WaitGroup) {
	// Read all files in directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() && recursive {
			countKindDir(directory+"/"+file.Name(), kind, recursive, wg)
		} else if file.IsDir() {
			continue
		}
		if utils.FileExtension(file.Name(), 2) != ".ast.json" {
			continue
		}
		wg.Add(1)
		go func(fileName, kind string) {
			defer wg.Done()
			countKindFile(fileName, kind)
		}(directory+"/"+file.Name(), kind)

	}
}

func countKindFile(fileName, kind string) {
	// Load file
	fileJSON, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var treeNode ast.Node
	err = json.Unmarshal(fileJSON, &treeNode)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON in file %s: %v\n", fileName, err)
		os.Exit(1)
	}
	v := &ast.VisitorCount{Kind: kind}
	treeNode.WalkPostfix(v)
	fmt.Printf("%s : Number of nodes of kind %s: %d\n", fileName, kind, v.Count)
}
