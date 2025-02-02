package cmd

import (
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/cmd/operations"
	"os"
)

const VERSION = "0.1.0"

func Main() {
	mainHelp := flag.Bool("help", false, "Show help for the program")
	mainVersion := flag.Bool("version", false, "Show the version of the program")

	flag.Parse()

	if *mainHelp {
		fmt.Println("Usage: go-php-parser [command] [flags]")
		fmt.Println("Commands:")
		fmt.Println("  parse - Parse a PHP file and output a JSON file with the tree")
		fmt.Println("  show - Show the tree of a JSON file")
		fmt.Println("  operations - Input a JSON tree file and then perform some operations on it")
		fmt.Println("Type ./main [command] --help for more information on a command")
		os.Exit(0)
	}

	if *mainVersion {
		fmt.Printf("Version: %s\n", VERSION)
		os.Exit(0)
	}

	// Check which command is being run
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please provide a command. Type --help for more information")
		os.Exit(1)
	}
	switch args[0] {
	case "parse":
		parsePHP(args)
	case "show":
		showTree(args)
	case "operations":
		operations.Main(args)
	default:
		fmt.Println("Please provide a valid command. Type --help for more information")
		os.Exit(1)
	}
}
