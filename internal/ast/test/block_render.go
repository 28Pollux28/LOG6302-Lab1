package test

import (
	"github.com/28Pollux28/log6302-parser/internal/ast"
	"github.com/28Pollux28/log6302-parser/utils"
)

func GetRenders() map[string]func(*utils.Stack, *ast.Node) IBlock {
	return map[string]func(*utils.Stack, *ast.Node) IBlock{
		"comment": func(s *utils.Stack, n *ast.Node) IBlock {
			return &PrimitiveBlock{Content: n.Text, BlockType: CommentBlockType}
		},
	}
}
