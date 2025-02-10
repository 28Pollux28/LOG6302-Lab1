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
		"literal": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: LiteralBlockType,
			}
		},
		"float": func(s *utils.Stack, n Node) IBlock {
			return &PrimitiveBlock{
				Content:   n.GetText(),
				BlockType: FloatBlockType,
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
			return &HorizontalBlock{
				Blocks:    blocks,
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
				}
			}
			// If there is a colon block, we need to put the endwhile and the semicolon in a horizontal block, then put them on a new line
			if colonBlockFlag {
				result = []IBlock{
					&HorizontalBlock{
						Blocks:    result[:len(result)-2],
						BlockType: COMPOSITE,
					},
					&HorizontalBlock{ //Put the endwhile and the semicolon on the same line
						Blocks:    result[len(result)-2:],
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
						result = append(result, block)
					}
					result = append(result, block, WHITESPACE_BLOCK)
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
				result = []IBlock{
					&HorizontalBlock{
						Blocks:    result,
						BlockType: COMPOSITE,
					},
					&IndentBlock{
						Block: &VerticalBlock{
							Blocks:    colonStatementBlocks[:len(colonStatementBlocks)-2],
							BlockType: COMPOSITE,
						},
					},
					&HorizontalBlock{ //Put the endfor and the semicolon on the same line
						Blocks:    colonStatementBlocks[len(colonStatementBlocks)-2:],
						BlockType: COMPOSITE,
					},
				}
				return &VerticalBlock{
					Blocks:    result,
					BlockType: ForStatementBlockType,
				}
			}
			return &HorizontalBlock{
				Blocks:    result,
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
		"assignement_expression": func(s *utils.Stack, n Node) IBlock {
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
		"variable_name": func(s *utils.Stack, n Node) IBlock {
			blocks := PopBlocksFromStack(s, n.GetChildrenNumber())
			return &HorizontalBlock{
				Blocks:    blocks,
				BlockType: VariableNameBlockType,
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
