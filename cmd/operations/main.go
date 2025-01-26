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
	operationsCmd.Parse(args[1:])

	if *help {
		fmt.Println("Usage: ./main operations [flags] file.json <operation> [operation_flags]")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the operations command")
		fmt.Println("Operations:")
		fmt.Println("  count-kind - Count the number of nodes of a specific kind")
		fmt.Println("  count-kinds - Count the number of nodes of multiple kinds")
		fmt.Println("  find-kind-tree - Find the tree of nodes of a specific kind")
		os.Exit(0)
	}

	if len(operationsCmd.Args()) < 2 {
		fmt.Println("Please provide a file name and an operation. Type --help for more information")
		os.Exit(1)
	}
	fileName := operationsCmd.Args()[0]
	operation := operationsCmd.Args()[1]
	switch operation {
	case "count-kind":
		countKind(fileName, operationsCmd.Args())
	case "count-kinds":
		countKinds(fileName, operationsCmd.Args())
	case "find-kind-tree":
		findKindTree(fileName, operationsCmd.Args())
	default:
		fmt.Println("Please provide a valid operation. Type --help for more information")
		os.Exit(1)
	}
}
