package operations

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/tree"
	"os"
)

func countKinds(fileName string, args []string) {
	countKindsOperation := flag.NewFlagSet("count-kinds", flag.ExitOnError)
	countKindsHelp := countKindsOperation.Bool("help", false, "Show help for the count-kind operation")
	countKindsOperation.Parse(args[2:])

	if *countKindsHelp {
		fmt.Println("Usage: ./main operations [flags] file.json count-kinds [flags] <kind1> <kind2> ...")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the count-kinds operation")
		fmt.Println(" <kind1> <kind2> ... - The kinds of nodes to count. Refer to the PHP tree-sitter grammar for the kinds")
		os.Exit(0)
	}

	if len(countKindsOperation.Args()) < 1 {
		fmt.Println("Please provide at least one kind. Type --help for more information")
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

	kinds := countKindsOperation.Args()
	counts := treeNode.CountKinds(kinds)
	for kind, count := range counts {
		fmt.Printf("%s: %d\n", kind, count)
	}
	os.Exit(0)
}
