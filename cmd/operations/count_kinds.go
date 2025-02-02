package operations

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/tree"
	"github.com/28Pollux28/log6302-parser/utils"
	"os"
	"sync"
)

func countKinds(fileName string, args []string, directory, recursive bool) {
	countKindsOperation := flag.NewFlagSet("count-kinds", flag.ExitOnError)
	countKindsHelp := countKindsOperation.Bool("help", false, "Show help for the count-kind operation")
	countKindsOperation.Parse(args[2:])

	if *countKindsHelp {
		fmt.Println("Usage: go-php-parser operations [OPFlags] <file.ast.json|directory> count-kinds [flags] <kind1> <kind2> ...")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the count-kinds operation")
		fmt.Println("  <kind1> <kind2> ... - The kinds of nodes to count. Refer to the PHP tree-sitter grammar for the kinds")
		os.Exit(0)
	}

	if len(countKindsOperation.Args()) < 1 {
		fmt.Println("Please provide at least one kind. Type --help for more information")
		os.Exit(1)
	}

	if directory {
		var wg sync.WaitGroup
		countKindsDir(fileName, countKindsOperation.Args(), recursive, &wg)
		wg.Wait()
		return
	}
	countKindsFile(fileName, countKindsOperation.Args())
}

func countKindsDir(directory string, kinds []string, recursive bool, wg *sync.WaitGroup) {
	// Read all files in directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() && recursive {
			countKindsDir(directory+"/"+file.Name(), kinds, recursive, wg)
		} else if file.IsDir() {
			continue
		}
		if utils.FileExtension(file.Name(), 2) != ".ast.json" {
			continue
		}
		wg.Add(1)
		go func(fileName string, kind []string) {
			defer wg.Done()
			countKindsFile(fileName, kinds)
		}(directory+"/"+file.Name(), kinds)

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
