package pretty_print

import "github.com/28Pollux28/log6302-parser/utils"

func PopBlocksFromStack(s *utils.Stack, n int) []IBlock {
	var blocks []IBlock
	for range n {
		blocks = append(blocks, s.Pop().(IBlock))
	}
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
