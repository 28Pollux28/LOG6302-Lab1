package pretty_print

import (
	"github.com/28Pollux28/log6302-parser/utils"
)

func GetRenders() map[string]func(*utils.Stack, Node) IBlock {
	return map[string]func(*utils.Stack, Node) IBlock{
		"program": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &VerticalBlock{
				Blocks:    blocks,
				BlockType: ProgramBlockType,
			}
		},
		"php_tag": func(s *utils.Stack, n Node) IBlock {
			return PHP_TAG_BLOCK
		},
		"text_interpolation": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &VerticalBlock{
				Blocks:    blocks,
				BlockType: TextInterpolationBlockType,
			}
		},
		"text": func(s *utils.Stack, n Node) IBlock {
			return &PrimitiveBlock{
				Content:   n.GetText(),
				BlockType: TextBlockType,
			}
		},
		"named_type": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: NamedTypeBlockType,
			}
		},
		"bottom_type": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: BottomTypeBlockType,
			}
		},
		"union_type": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: UnionTypeBlockType,
			}
		},
		"intersection_type": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: IntersectionTypeBlockType,
			}
		},
		"primitive_type": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: PrimitiveTypeBlockType,
			}
		},
		"cast_type": func(s *utils.Stack, n Node) IBlock {
			return &PrimitiveBlock{
				Content:   n.GetText(),
				BlockType: CastTypeBlockType,
			}
		},
		"const_element": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, WHITESPACE_BLOCK),
				BlockType: ConstElementBlockType,
			}
		},
		"echo_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			finalBlocks := []IBlock{blocks[0]}
			finalBlocks = append(finalBlocks, WHITESPACE_BLOCK)
			finalBlocks = append(finalBlocks, blocks[1:]...)
			return &HorizontalBlock{
				Blocks:    finalBlocks,
				BlockType: EchoStatementBlockType,
			}
		},
		"exit_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: ExitStatementBlockType,
			}
		},
		"unset_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: UnsetStatementBlockType,
			}
		},
		"float": func(s *utils.Stack, n Node) IBlock {
			return &PrimitiveBlock{
				Content:   n.GetText(),
				BlockType: FloatBlockType,
			}
		},
		"try_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			for _, block := range blocks {
				switch block.Type() {
				case TryBlockType:
					result = append(result, block, WHITESPACE_BLOCK)
				default:
					result = append(result, block)
				}
			}
			return &HorizontalBlock{
				Blocks:    append(result, NEWLINE_BLOCK),
				BlockType: TryStatementBlockType,
			}
		},
		"catch_clause": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			for i, block := range blocks {
				switch block.Type() {
				case CatchBlockType:
					result = append(result, WHITESPACE_BLOCK, block, WHITESPACE_BLOCK)
				case CloseParenBlockType:
					result = append(result, block)
					if i != len(blocks)-1 { //We need to look ahead to put the whitespace
						result = append(result, WHITESPACE_BLOCK)
					}
				case TypeListBlockType:
					result = append(result, block, WHITESPACE_BLOCK)
				default:
					result = append(result, block)
				}
			}
			return &HorizontalBlock{
				Blocks:    append(result, NEWLINE_BLOCK),
				BlockType: CatchClauseBlockType,
			}
		},
		"type_list": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, WHITESPACE_BLOCK),
				BlockType: TypeListBlockType,
			}
		},
		"finally_clause": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			for _, block := range blocks {
				switch block.Type() {
				case FinallyBlockType:
					result = append(result, block, WHITESPACE_BLOCK)
				default:
					result = append(result, block)
				}
			}
			return &HorizontalBlock{
				Blocks:    append(result, NEWLINE_BLOCK),
				BlockType: FinallyClauseBlockType,
			}
		},
		"goto_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			finalBlocks := []IBlock{blocks[0]}
			finalBlocks = append(finalBlocks, WHITESPACE_BLOCK)
			finalBlocks = append(finalBlocks, blocks[1:]...)
			return &HorizontalBlock{
				Blocks:    finalBlocks,
				BlockType: GotoStatementBlockType,
			}
		},
		"continue_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			finalBlocks := []IBlock{blocks[0]}
			finalBlocks = append(finalBlocks, WHITESPACE_BLOCK)
			finalBlocks = append(finalBlocks, blocks[1:]...)
			return &HorizontalBlock{
				Blocks:    finalBlocks,
				BlockType: ContinueStatementBlockType,
			}
		},
		"break_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			finalBlocks := []IBlock{blocks[0]}
			if len(blocks) > 2 {
				finalBlocks = append(finalBlocks, WHITESPACE_BLOCK)
			}
			finalBlocks = append(finalBlocks, blocks[1:]...)
			return &HorizontalBlock{
				Blocks:    finalBlocks,
				BlockType: BreakStatementBlockType,
			}
		},
		"integer": func(s *utils.Stack, n Node) IBlock {
			return &PrimitiveBlock{
				Content:   n.GetText(),
				BlockType: IntegerBlockType,
			}
		},
		"return_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: ReturnStatementBlockType,
			}
		},
		"throw_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			for _, block := range blocks {
				if block.Type() == ThrowBlockType {
					result = append(result, block, WHITESPACE_BLOCK)
				} else {
					result = append(result, block)
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: ThrowExpressionBlockType,
			}
		},
		"while_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			var colonBlockFlag bool
			for i, block := range blocks {
				switch block.Type() {
				case WhileBlockType:
					result = append(result, block, WHITESPACE_BLOCK)
				case ParenthesizedExpressionBlockType:
					result = append(result, block)
					if i != len(blocks)-1 && blocks[i+1].Type() != ColonBlockBlockType { //We need to look ahead to put the whitespace between the parenthesized expression and the statement {...}
						result = append(result, WHITESPACE_BLOCK)
					}
				case ColonBlockBlockType:
					colonBlockFlag = true
					result = append(result, block)
				case EndwhileBlockType, SemicolonBlockType:
					result = append(result, block)
				default:
					result = append(result, block)
				}
			}
			result = append(result, NEWLINE_BLOCK)
			// If there is a colon block, we need to put the endwhile and the semicolon in a horizontal block, then put them on a new line
			if colonBlockFlag {
				result = []IBlock{
					&HorizontalBlock{
						Blocks:    result[:len(result)-3],
						BlockType: COMPOSITE,
					},
					&HorizontalBlock{ //Put the endwhile and the semicolon on the same line
						Blocks:    result[len(result)-3:],
						BlockType: COMPOSITE,
					},
				}
				return &VerticalBlock{
					Blocks:    result,
					BlockType: WhileStatementBlockType,
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: WhileStatementBlockType,
			}
		},
		"do_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: DoStatementBlockType,
			}
		},
		"for_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			var colonBlockFlag bool
			var colonStatementBlocks []IBlock
			for i, block := range blocks {
				switch block.Type() {
				case ForBlockType:
					result = append(result, block, WHITESPACE_BLOCK)
				case OpenParenBlockType:
					result = append(result, block)
				case CloseParenBlockType:
					result = append(result, block)
					if i != len(blocks)-1 && (blocks[i+1].Type() != ColonBlockType && blocks[i+1].Type() != SemicolonBlockType) { //We need to look ahead to put the whitespace between the parenthesized expression and the statement {...}
						result = append(result, WHITESPACE_BLOCK)
					}
				case SemicolonBlockType:
					if colonBlockFlag {
						colonStatementBlocks = append(colonStatementBlocks, block)
					} else {
						result = append(result, block, WHITESPACE_BLOCK)
					}
				case ColonBlockType:
					colonBlockFlag = true
					result = append(result, block)
				case EndforBlockType:
					colonStatementBlocks = append(colonStatementBlocks, block)
				default:
					if colonBlockFlag {
						colonStatementBlocks = append(colonStatementBlocks, block)
					} else {
						result = append(result, block)
					}
				}
			}
			// If there is a colon block, we need to put the endfor and the semicolon in a horizontal block, then put them on a new line
			if colonBlockFlag {
				colonStatementBlocks = append(colonStatementBlocks, NEWLINE_BLOCK)
				result = []IBlock{
					&HorizontalBlock{
						Blocks:    result,
						BlockType: COMPOSITE,
					},
					&IndentBlock{
						Block: &VerticalBlock{
							Blocks:      colonStatementBlocks[:len(colonStatementBlocks)-3],
							BlockType:   COMPOSITE,
							IndentFirst: true,
						},
					},
					&HorizontalBlock{ //Put the endfor and the semicolon on the same line
						Blocks:    colonStatementBlocks[len(colonStatementBlocks)-3:],
						BlockType: COMPOSITE,
					},
				}
				return &VerticalBlock{
					Blocks:    result,
					BlockType: ForStatementBlockType,
				}
			}
			return &HorizontalBlock{
				Blocks:    append(result, NEWLINE_BLOCK),
				BlockType: ForStatementBlockType,
			}
		},
		"sequence_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			for _, block := range blocks {
				if block.Type() == CommaBlockType {
					result = append(result, block, WHITESPACE_BLOCK)
				} else {
					result = append(result, block)
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: SequenceExpressionBlockType,
			}
		},
		//"foreach_statement": func(s *utils.Stack, n Node) IBlock {
		//	blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
		//	var result []IBlock
		//	var colonBlockFlag bool
		//	for i, block := range blocks {
		//		blockType := block.Type()
		//		switch {
		//		case isBlockOfType(blockType, ForeachBlockType):
		//			result = append(result, block, WHITESPACE_BLOCK)
		//		case isBlockOfType(blockType, OpenParenBlockType):
		//			result = append(result, block)
		//		case isBlockOfType(blockType, AsBlockType):
		//			result = append(result, WHITESPACE_BLOCK, block, WHITESPACE_BLOCK)
		//		case isBlockOfType(blockType, CloseParenBlockType):
		//			result = append(result, block)
		//			if i != len(blocks)-1 && (blocks[i+1].Type() != ColonBlockType && blocks[i+1].Type() != SemicolonBlockType) { //We need to look ahead to put the whitespace between the parenthesized expression and the statement {...}
		//				result = append(result, WHITESPACE_BLOCK)
		//			}
		//		case isBlockOfType(blockType, ColonBlockBlockType):
		//			colonBlockFlag = true
		//			result = append(result, block)
		//		case isBlockOfType(blockType, EndforeachBlockType):
		//			result = append(result, block)
		//		default:
		//			result = append(result, block)
		//		}
		//	}
		//},
		"if_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			var colonBlockFlag bool
			for i, block := range blocks {
				switch block.Type() {
				case IfBlockType:
					result = append(result, block, WHITESPACE_BLOCK)
				case ParenthesizedExpressionBlockType:
					result = append(result, block)
					if i != len(blocks)-1 && blocks[i+1].Type() != ColonBlockBlockType { //We need to look ahead to put the whitespace between the parenthesized expression and the statement {...}
						result = append(result, WHITESPACE_BLOCK)
					}
				case ColonBlockBlockType:
					colonBlockFlag = true
					result = append(result, block)
				case EndifBlockType, SemicolonBlockType:
					result = append(result, block)
				case ElseIfClauseBlockType, ElseClauseBlockType:
					if !colonBlockFlag {
						result = append(result, WHITESPACE_BLOCK, block)
					} else {
						result = append(result, NEWLINE_BLOCK, block)
					}

				default:
					result = append(result, block)
				}
			}
			result = append(result, NEWLINE_BLOCK)
			// If there is a colon block, we need to put the endif and the semicolon in a horizontal block, then put them on a new line
			if colonBlockFlag {
				result = []IBlock{
					&HorizontalBlock{
						Blocks:    result[:len(result)-3],
						BlockType: COMPOSITE,
					},
					&HorizontalBlock{ //Put the endif and the semicolon on the same line
						Blocks:    result[len(result)-3:],
						BlockType: COMPOSITE,
					},
				}
				return &VerticalBlock{
					Blocks:    result,
					BlockType: IfStatementBlockType,
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: IfStatementBlockType,
			}
		},
		"colon_block": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks: []IBlock{
					&IndentBlock{
						Block: &VerticalBlock{
							Blocks:    blocks,
							BlockType: COMPOSITE,
						},
					},
				},
				BlockType: ColonBlockBlockType,
			}
		},
		"else_if_clause": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			var colonBlockFlag bool
			for i, block := range blocks {
				switch block.Type() {
				case ElseifBlockType:
					result = append(result, block, WHITESPACE_BLOCK)
				case ParenthesizedExpressionBlockType:
					result = append(result, block)
					if i != len(blocks)-1 && blocks[i+1].Type() != ColonBlockBlockType { //We need to look ahead to put the whitespace between the parenthesized expression and the statement {...}
						result = append(result, WHITESPACE_BLOCK)
					}
				case ColonBlockBlockType:
					colonBlockFlag = true
					result = append(result, block)
				case SemicolonBlockType:
					result = append(result, block)
				default:
					result = append(result, block)
				}
			}
			if colonBlockFlag {
				result = []IBlock{
					&HorizontalBlock{
						Blocks:    result,
						BlockType: COMPOSITE,
					},
				}
				return &VerticalBlock{
					Blocks:    result,
					BlockType: ElseIfClauseBlockType,
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: ElseIfClauseBlockType,
			}
		},
		"else_clause": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			var colonBlockFlag bool
			for i, block := range blocks {
				switch block.Type() {
				case ElseBlockType:
					if i != len(blocks)-1 && blocks[i+1].Type() != ColonBlockBlockType { //We need to look ahead to put the whitespace between the else and the statement {...}
						result = append(result, block, WHITESPACE_BLOCK)
					} else {
						result = append(result, block)
					}
				case ColonBlockBlockType:
					colonBlockFlag = true
					result = append(result, block)
				case SemicolonBlockType:
					result = append(result, block)
				default:
					result = append(result, block)
				}
			}
			if colonBlockFlag {
				result = []IBlock{
					&HorizontalBlock{
						Blocks:    result,
						BlockType: COMPOSITE,
					},
				}
				return &VerticalBlock{
					Blocks:    result,
					BlockType: ElseClauseBlockType,
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: ElseClauseBlockType,
			}
		},
		"pair": func(s *utils.Stack, n Node) IBlock { //Actually foreach_pair in the grammar, but under alias for some reason
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, WHITESPACE_BLOCK),
				BlockType: ForeachPairBlockType,
			}
		},
		"switch_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks: []IBlock{
					&HorizontalBlock{
						Blocks:    append(blocks, NEWLINE_BLOCK),
						BlockType: COMPOSITE,
					},
				},
				BlockType: SwitchStatementBlockType,
			}
		},
		"switch_block": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			var colonBlockFlag bool
			for _, block := range blocks {
				switch {
				case isBlockOfType(block.Type(), ColonBlockType):
					colonBlockFlag = true
					result = append(result, block)
				default:
					result = append(result, block)
				}
			}
			if colonBlockFlag {
				finalBlock := []IBlock{result[0]}
				finalBlock = append(finalBlock,
					&IndentBlock{
						Block: &VerticalBlock{
							Blocks:      result[1 : len(result)-2],
							BlockType:   COMPOSITE,
							IndentFirst: true,
						},
					}, &HorizontalBlock{
						Blocks:    result[len(result)-2:],
						BlockType: COMPOSITE,
					},
					&HorizontalBlock{
						Blocks: []IBlock{NEWLINE_BLOCK},
					})

				return &VerticalBlock{
					Blocks:    finalBlock,
					BlockType: SwitchBlockBlockType,
				}
			}
			finalBlock := []IBlock{&HorizontalBlock{Blocks: []IBlock{WHITESPACE_BLOCK, result[0]}, BlockType: COMPOSITE}}
			finalBlock = append(finalBlock, &IndentBlock{
				Block: &VerticalBlock{
					Blocks:      result[1 : len(result)-1],
					BlockType:   COMPOSITE,
					IndentFirst: true,
				},
			}, result[len(result)-1])

			return &VerticalBlock{
				Blocks:    finalBlock,
				BlockType: SwitchBlockBlockType,
			}
		},
		"case_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			var colonBlockFlag bool
			var colonStatementBlocks []IBlock
			for i, block := range blocks {
				switch {
				case isBlockOfType(block.Type(), CaseBlockType):
					result = append(result, block, WHITESPACE_BLOCK)
				case isExpressionBlockType(block.Type()):
					result = append(result, block)
				case isBlockOfType(block.Type(), SemicolonBlockType):
					if i == len(blocks)-1 {
						result = append(result, block)
					} else {
						colonBlockFlag = true
						colonStatementBlocks = append(colonStatementBlocks, block)
					}
				case isBlockOfType(block.Type(), ColonBlockType):
					colonBlockFlag = true
					colonStatementBlocks = append(colonStatementBlocks, block)
				case isStatementBlockType(block.Type()):
					if colonBlockFlag {
						colonStatementBlocks = append(colonStatementBlocks, block)
					}
					// This should never happen
				}
			}
			if colonBlockFlag {
				result = []IBlock{
					&HorizontalBlock{
						Blocks:    result,
						BlockType: COMPOSITE,
					},
					&IndentBlock{
						Block: &VerticalBlock{
							Blocks:    colonStatementBlocks,
							BlockType: COMPOSITE,
						},
					},
				}
				return &HorizontalBlock{
					Blocks:    result,
					BlockType: CaseStatementBlockType,
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: CaseStatementBlockType,
			}
		},
		"default_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			var colonBlockFlag bool
			var colonStatementBlocks []IBlock
			for i, block := range blocks {
				switch {
				case isBlockOfType(block.Type(), DefaultBlockType):
					result = append(result, block)
				case isBlockOfType(block.Type(), SemicolonBlockType):
					if i == len(blocks)-1 {
						result = append(result, block)
					} else {
						colonBlockFlag = true
						colonStatementBlocks = append(colonStatementBlocks, block)
					}
				case isBlockOfType(block.Type(), ColonBlockType):
					colonBlockFlag = true
					colonStatementBlocks = append(colonStatementBlocks, block)
				case isStatementBlockType(block.Type()):
					if colonBlockFlag {
						colonStatementBlocks = append(colonStatementBlocks, block)
					}
					// This should never happen
				}
			}
			if colonBlockFlag {
				result = []IBlock{
					&HorizontalBlock{
						Blocks:    result,
						BlockType: COMPOSITE,
					},
					&IndentBlock{
						Block: &VerticalBlock{
							Blocks:    colonStatementBlocks,
							BlockType: COMPOSITE,
						},
					},
				}
				return &HorizontalBlock{
					Blocks:    result,
					BlockType: DefaultStatementBlockType,
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: DefaultStatementBlockType,
			}
		},
		"compound_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &VerticalBlock{
				Blocks: []IBlock{
					blocks[0],
					&IndentBlock{
						Block: &VerticalBlock{
							Blocks:      blocks[1 : len(blocks)-1],
							BlockType:   COMPOSITE,
							IndentFirst: true,
						},
					},
					blocks[len(blocks)-1],
				},
				BlockType: CompoundStatementBlockType,
			}
		},
		"named_label_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: NamedLabelStatementBlockType,
			}
		},
		"expression_statement": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: ExpressionStatementBlockType,
			}
		},
		"unary_op_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: UnaryOpExpressionBlockType,
			}
		},
		"error_suppression_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: ErrorSuppressionExpressionBlockType,
			}
		},
		"clone_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    []IBlock{blocks[0], WHITESPACE_BLOCK, blocks[1]},
				BlockType: CloneExpressionBlockType,
			}
		},
		"parenthesized_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: ParenthesizedExpressionBlockType,
			}
		},
		"print_intrinsic": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, WHITESPACE_BLOCK),
				BlockType: PrintIntrinsicBlockType,
			}
		},
		"object_creation_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, WHITESPACE_BLOCK),
				BlockType: ObjectCreationExpressionBlockType,
			}
		},
		"update_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: UpdateExpressionBlockType,
			}
		},
		"cast_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			for i, block := range blocks {
				result = append(result, block)
				if i == len(blocks)-2 {
					result = append(result, WHITESPACE_BLOCK)
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: CastExpressionBlockType,
			}
		},
		"cast_variable": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			for i, block := range blocks {
				result = append(result, block)
				if i == len(blocks)-2 {
					result = append(result, WHITESPACE_BLOCK)
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: CastVariableBlockType,
			}
		},
		"assignment_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, WHITESPACE_BLOCK),
				BlockType: AssignmentExpressionBlockType,
			}
		},
		"reference_assignment_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			var result []IBlock
			for i, block := range blocks {
				result = append(result, block)
				if i != len(blocks)-1 || block.Type() == ReferenceModifierBlockType {
					result = append(result, WHITESPACE_BLOCK)
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
				BlockType: ReferenceAssignmentExpressionBlockType,
			}
		},
		"conditional_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, WHITESPACE_BLOCK),
				BlockType: ConditionalExpressionBlockType,
			}
		},
		"augmented_assignment_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, WHITESPACE_BLOCK),
				BlockType: AugmentedAssignmentExpressionBlockType,
			}
		},
		"arguments": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, &HorizontalBlock{[]IBlock{COMMA_BLOCK, WHITESPACE_BLOCK}, COMPOSITE}),
				BlockType: ArgumentsBlockType,
			}
		},
		"argument": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: ArgumentBlockType,
			}
		},
		"member_call_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: MemberCallExpressionBlockType,
			}
		},
		"escape_sequence": func(s *utils.Stack, n Node) IBlock {
			return &PrimitiveBlock{
				Content:   n.GetText(),
				BlockType: EscapeSequenceBlockType,
			}
		},
		"encapsed_string": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: EncapsedStringBlockType,
			}
		},
		"string": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, WHITESPACE_BLOCK),
				BlockType: StringBlockType,
			}
		},
		"string_content": func(s *utils.Stack, n Node) IBlock {
			return &PrimitiveBlock{
				Content:   n.GetText(),
				BlockType: StringContentBlockType,
			}
		},
		"variable_name": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: VariableNameBlockType,
			}
		},
		"by_ref": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: ByRefBlockType,
			}
		},
		"binary_expression": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    joinBlocks(blocks, WHITESPACE_BLOCK),
				BlockType: BinaryExpressionBlockType,
			}
		},
		"name": func(s *utils.Stack, n Node) IBlock {
			return &PrimitiveBlock{
				Content:   n.GetText(),
				BlockType: NameBlockType,
			}
		},
		"comment": func(s *utils.Stack, n Node) IBlock {
			return &PrimitiveBlock{
				Content:   n.GetText(),
				BlockType: CommentBlockType,
			}
		},
	}
}
