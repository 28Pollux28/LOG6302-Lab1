package ast

import (
	"fmt"
	"github.com/28Pollux28/log6302-parser/internal/ast/pretty_print"
	"github.com/28Pollux28/log6302-parser/utils"
	"maps"
)

type PrettyPrintVisitor struct {
	Result     utils.Stack
	NodeConfig map[string]func(*utils.Stack, pretty_print.Node) pretty_print.IBlock
}

func NewPrettyPrintVisitor() *PrettyPrintVisitor {
	return &PrettyPrintVisitor{
		Result:     utils.Stack{},
		NodeConfig: initNodeConfig(),
	}
}

func (p *PrettyPrintVisitor) VisitNode(n *Node) {
	blockFunc, ok := p.NodeConfig[n.Kind]
	if !ok {
		fmt.Printf("\033[0;31m No block for node kind %s \033[0m\n", n.Kind)
		blockFunc = func(stack *utils.Stack, node pretty_print.Node) pretty_print.IBlock {
			blocks := pretty_print.PopBlocksFromStack(stack, node.GetChildrenNumber())
			return &pretty_print.HorizontalBlock{
				Blocks: blocks,
			}
		}
	}
	p.Result.Push(blockFunc(&p.Result, n))
}

// initNodeConfig initializes the node configuration for the pretty print visitor
// It returns a map of node kinds to functions that return a block for that node
func initNodeConfig() map[string]func(*utils.Stack, pretty_print.Node) pretty_print.IBlock {
	primitiveBlockMap := pretty_print.GetPrimitiveBlockRenders()
	blockRendersMap := pretty_print.GetRenders()
	maps.Copy(blockRendersMap, primitiveBlockMap)
	return blockRendersMap
}

func (p *PrettyPrintVisitor) Print() string {
	block := p.Result.Pop().(pretty_print.IBlock)
	return block.Render(0)
}
