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

func createCfg(fileName string, args []string, directory, recursive bool) {
	createCfgOperation := flag.NewFlagSet("create-cfg", flag.ExitOnError)
	outputFile := createCfgOperation.String("output", "", "The output JSON file")
	createCfgHelp := createCfgOperation.Bool("help", false, "Show help for the create-Cfg operation")
	createCfgOperation.Parse(args[2:])

	if *createCfgHelp {
		fmt.Println("Usage: go-php-parser operations <file.ast.json|directory> create-Cfg [flags]")
		fmt.Println("Flags:")
		fmt.Println("  --output - The output JSON file / directory")
		fmt.Println("  --help - Show help for the create-Cfg operation")
		os.Exit(0)
	}

	if len(createCfgOperation.Args()) > 1 {
		fmt.Println("Too many arguments. Type --help for more information")
		os.Exit(1)
	}

	if directory {
		var wg sync.WaitGroup
		createCfgDir(fileName, *outputFile, recursive, &wg)
		wg.Wait()
		return
	}
	createCfgFile(fileName, *outputFile)
	return
}

func createCfgDir(directory, outputFile string, recursive bool, wg *sync.WaitGroup) {
	// Read all files in directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() && recursive {
			createCfgDir(directory+"/"+file.Name(), outputFile+"/"+file.Name(), recursive, wg)
		} else if file.IsDir() {
			continue
		}
		if utils.FileExtension(file.Name(), 2) != ".ast.json" {
			continue
		}
		wg.Add(1)
		go func(fileName, outputFile string) {
			defer wg.Done()
			createCfgFile(fileName, outputFile)
		}(directory+"/"+file.Name(), outputFile+"/"+file.Name())

	}
}

func createCfgFile(fileName, outputFile string) {
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

	cfg := ast.BuildCFGFromAST(&treeNode)
	if outputFile != "" {
		outputFile = strings.TrimSuffix(outputFile, "php.ast.json")
		if !strings.HasSuffix(outputFile, ".dot") {
			outputFile += ".dot"
		}
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
		_, err = file.WriteString(cfg.GenerateDOT())
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}

	} else {
		fmt.Println(cfg.GenerateDOT())

	}

}
