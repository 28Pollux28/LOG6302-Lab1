package operations

import (
	"flag"
	"fmt"
	"os"
)

func Main(args []string) {
	var operationsCmd = flag.NewFlagSet("operations", flag.ExitOnError)
	// Define the flags for the operation command
	help := operationsCmd.Bool("help", false, "Show help for the operations command")
	directory := operationsCmd.Bool("directory", false, "Perform the operation on a directory of AST trees")
	recursive := operationsCmd.Bool("recursive", false, "Recursively perform the operation on a directory of AST trees")
	operationsCmd.Parse(args[1:])

	if *help {
		fmt.Println("Usage: go-php-parser operations [flags] file.ast.json <operation> [operation_flags]")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the operations command")
		fmt.Println("  --directory - Perform the operation on a directory of AST trees")
		fmt.Println("  --recursive - Recursively perform the operation on a directory of AST trees")
		fmt.Println("Operations:")
		fmt.Println("  count-kind - Count the number of nodes of a specific kind")
		fmt.Println("  count-kinds - Count the number of nodes of multiple kinds")
		fmt.Println("  find-kind-tree - Find the tree of nodes of a specific kind")
		fmt.Println("  find-kind-trees - Find the trees of nodes of a specific kind")
		os.Exit(0)
	}

	if len(operationsCmd.Args()) < 2 {
		fmt.Println("Please provide a file name and an operation. Type --help for more information")
		os.Exit(1)
	}
	if recursive != nil && *recursive && directory != nil && !*directory {
		fmt.Println("The --recursive flag can only be used with the --directory flag")
		os.Exit(1)
	}

	fileName := operationsCmd.Args()[0]
	operation := operationsCmd.Args()[1]
	switch operation {
	case "count-kind":
		countKind(fileName, operationsCmd.Args(), *directory, *recursive)
	case "count-kinds":
		countKinds(fileName, operationsCmd.Args(), *directory, *recursive)
	case "find-kind-tree":
		findKindTree(fileName, operationsCmd.Args(), *directory, *recursive)
	case "find-kind-trees":
		findKindTrees(fileName, operationsCmd.Args(), *directory, *recursive)
	case "pretty-print":
		prettyPrint(fileName, operationsCmd.Args(), *directory, *recursive)
	default:
		fmt.Println("Please provide a valid operation. Type --help for more information")
		os.Exit(1)
	}
}
