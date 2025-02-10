package ast

import (
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/ast/pretty_print"
	"github.com/28Pollux28/log6302-parser/utils"
	"maps"
)

type PrettyPrintVisitor struct {
	Result     utils.Stack
	NodeConfig map[string]func(*utils.Stack, *pretty_print.Node) pretty_print.IBlock
}

func (p *PrettyPrintVisitor) VisitNode(n *Node) {
	blockFunc, ok := p.NodeConfig[n.Kind]
	if !ok {
		fmt.Printf("\033[0;31m No block for node kind %s \033[0m\n", n.Kind)
		blockFunc = func(stack *utils.Stack, node *pretty_print.Node) pretty_print.IBlock {
			blocks := pretty_print.PopBlocksFromStack(stack, node.GetChildrenNumber())
			return &pretty_print.HorizontalBlock{
				Blocks: blocks,
			}
		}
	}
	p.Result.Push(blockFunc(&p.Result, &pretty_print.Node{
		Kind:      n.Kind,
		NChildren: n.GetChildrenNumber(),
		Text:      n.Text,
	}))
}

// initNodeConfig initializes the node configuration for the pretty print visitor
// It returns a map of node kinds to functions that return a block for that node
func (p *PrettyPrintVisitor) initNodeConfig() map[string]func(*utils.Stack, *pretty_print.Node) pretty_print.IBlock {
	primitiveBlockMap := pretty_print.GetPrimitiveBlockRenders()
	blockRendersMap := pretty_print.GetRenders()
	maps.Copy(primitiveBlockMap, blockRendersMap)
	return blockRendersMap
}
