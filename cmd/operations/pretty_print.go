package operations

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/28Pollux28/log6302-parser/internal/ast"
	"github.com/28Pollux28/log6302-parser/utils"
)

func prettyPrint(fileName string, args []string, directory, recursive bool) {
	prettyPrintOperation := flag.NewFlagSet("pretty-print", flag.ExitOnError)
	prettyPrintErrorOnly := prettyPrintOperation.Bool("error-only", false, "Only print errors")
	outputFile := prettyPrintOperation.String("output", "", "The output JSON file")
	prettyPrintHelp := prettyPrintOperation.Bool("help", false, "Show help for the pretty-print operation")
	prettyPrintOperation.Parse(args[2:])

	if *prettyPrintHelp {
		fmt.Println("Usage: go-php-parser operations <file.ast.json|directory> pretty-print [flags]")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the pretty-print operation")
		fmt.Println("  --output - The output PHP file / directory")
		fmt.Println("  --error-only - Only print errors - For debugging purposes")
		os.Exit(0)
	}

	if directory {
		var wg sync.WaitGroup
		prettyPrintDir(fileName, *outputFile, recursive, *prettyPrintErrorOnly, &wg)
		wg.Wait()
		return
	}
	prettyPrintFile(fileName, *outputFile, *prettyPrintErrorOnly)
	return
}

func prettyPrintDir(directory, output string, recursive, errorOnly bool, wg *sync.WaitGroup) {
	// Read all files in directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() && recursive {
			prettyPrintDir(directory+"/"+file.Name(), output+"/"+file.Name(), recursive, errorOnly, wg)
		} else if file.IsDir() {
			continue
		}
		if utils.FileExtension(file.Name(), 2) != ".ast.json" {
			continue
		}
		wg.Add(1)
		go func(fileName, outputFile string, errorOnly bool) {
			defer wg.Done()
			prettyPrintFile(fileName, outputFile, errorOnly)
		}(directory+"/"+file.Name(), output+"/"+file.Name(), errorOnly)

	}
}

func prettyPrintFile(fileName, outputFile string, errorOnly bool) {
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
	if !errorOnly && outputFile == "" {
		fmt.Printf("%s :\n%s\n", fileName, v.Print())
		return
	}
	if outputFile != "" {
		outputFile = strings.TrimSuffix(outputFile, ".ast.json")
		dir := path.Dir(outputFile)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0666)
			if err != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				os.Exit(1)
			}
		}
		file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		_, err = file.WriteString(v.Print())
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}
	}
}
