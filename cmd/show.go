package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/tree"
	"os"
)

func showTree(args []string) {
	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	showHelp := showCmd.Bool("help", false, "Show help for the show command")
	showCmd.Parse(args[1:])

	if *showHelp {
		fmt.Println("Usage: ./main show [flags] file.json")
		fmt.Println("Flags:")
		fmt.Println("  --help - Show help for the show command") //TODO: Add flags to specify which fields to show in the tree
		os.Exit(0)
	}

	if len(showCmd.Args()) < 1 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	// Load file name from args
	fileName := showCmd.Args()[0]
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var treeNode tree.Node
	err = json.Unmarshal(file, &treeNode)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	treeNode.PrintTree()
	os.Exit(0)
}
