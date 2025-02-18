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

func createCfg(fileName string, args []string, directory, recursive bool) {
	createCfgOperation := flag.NewFlagSet("create-cfg", flag.ExitOnError)
	createCfgHelp := createCfgOperation.Bool("help", false, "Show help for the create-Cfg operation")
	createCfgOperation.Parse(args[2:])

	if *createCfgHelp {
		fmt.Println("Usage: go-php-parser operations <file.ast.json|directory> create-Cfg [flags]")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the create-Cfg operation")
		os.Exit(0)
	}

	if len(createCfgOperation.Args()) > 10 {
		fmt.Println("Too many arguments. Type --help for more information")
		os.Exit(1)
	}

	if directory {
		var wg sync.WaitGroup
		createCfgDir(fileName, recursive, &wg)
		wg.Wait()
		return
	}
	createCfgFile(fileName)
	return
}

func createCfgDir(directory string, recursive bool, wg *sync.WaitGroup) {
	// Read all files in directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() && recursive {
			createCfgDir(directory+"/"+file.Name(), recursive, wg)
		} else if file.IsDir() {
			continue
		}
		if utils.FileExtension(file.Name(), 2) != ".ast.json" {
			continue
		}
		wg.Add(1)
		go func(fileName, kind string) {
			defer wg.Done()
			createCfgFile(fileName)
		}(directory+"/"+file.Name(), "")

	}
}

func createCfgFile(fileName string) {
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
	fmt.Println(cfg.GenerateDOT())

}
