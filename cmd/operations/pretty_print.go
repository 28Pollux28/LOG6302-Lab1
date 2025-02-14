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

func prettyPrint(fileName string, args []string, directory, recursive bool) {
	prettyPrintOperation := flag.NewFlagSet("pretty-print", flag.ExitOnError)
	prettyPrintHelp := prettyPrintOperation.Bool("help", false, "Show help for the pretty-print operation")
	prettyPrintOperation.Parse(args[2:])
	//TODO: add option to pretty print directly to a file

	if *prettyPrintHelp {
		fmt.Println("Usage: go-php-parser operations <file.ast.json|directory> pretty-print [flags]")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the pretty-print operation")
		os.Exit(0)
	}

	if directory {
		var wg sync.WaitGroup
		prettyPrintDir(fileName, recursive, &wg)
		wg.Wait()
		return
	}
	prettyPrintFile(fileName)
	return
}

func prettyPrintDir(directory string, recursive bool, wg *sync.WaitGroup) {
	// Read all files in directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() && recursive {
			prettyPrintDir(directory+"/"+file.Name(), recursive, wg)
		} else if file.IsDir() {
			continue
		}
		if utils.FileExtension(file.Name(), 2) != ".ast.json" {
			continue
		}
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			prettyPrintFile(fileName)
		}(directory + "/" + file.Name())

	}
}

func prettyPrintFile(fileName string) {
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
	v := ast.NewPrettyPrintVisitor()
	treeNode.WalkPostfix(v)
	fmt.Printf("%s :\n%s\n", fileName, v.Print())
}
