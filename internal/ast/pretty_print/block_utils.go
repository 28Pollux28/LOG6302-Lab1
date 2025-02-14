package pretty_print

import (
	"fmt"
	"github.com/28Pollux28/log6302-parser/utils"
	"slices"
)

func PopBlocksFromStack(s *utils.Stack, n int) []IBlock {
	var blocks []IBlock
	for range n {
		pop := s.Pop()
		if pop == nil {
			fmt.Println("Error: Stack is empty")
		}
		blocks = append(blocks, pop.(IBlock))
	}
	slices.Reverse(blocks)
	return blocks
}

func joinBlocks(blocks []IBlock, separator IBlock) []IBlock {
	var result []IBlock
	for i, block := range blocks {
		result = append(result, block)
		if i != len(blocks)-1 {
			result = append(result, separator)
		}
	}
	return result
}
