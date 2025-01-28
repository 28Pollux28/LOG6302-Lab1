package operations

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/tree"
	"os"
)

func countKind(fileName string, args []string) {
	countKindOperation := flag.NewFlagSet("count-kind", flag.ExitOnError)
	countKindHelp := countKindOperation.Bool("help", false, "Show help for the count-kind operation")
	countKindOperation.Parse(args[2:])

	if *countKindHelp {
		fmt.Println("Usage: ./main operations [flags] file.json count-kind [flags] <kind>")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the count-kind operation")
		fmt.Println("  <kind> - The kind of node to count. Refer to the PHP tree-sitter grammar for the kinds")
		os.Exit(0)
	}

	if len(countKindOperation.Args()) < 1 {
		fmt.Println("Please provide a kind. Type --help for more information")
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
		fmt.Println(err)
		os.Exit(1)
	}

	kind := countKindOperation.Args()[0]
	count := treeNode.CountKind(kind)
	fmt.Printf("Number of nodes of kind %s: %d\n", kind, count)
	os.Exit(0)
}
