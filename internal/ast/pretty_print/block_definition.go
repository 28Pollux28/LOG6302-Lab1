package pretty_print

import (
	"bytes"
	"strings"
)

type IBlock interface {
	Render(indentLvl int) string
	Type() BlockType
}

type HorizontalBlock struct {
	Blocks    []IBlock
	BlockType BlockType
}

func (h *HorizontalBlock) Render(indentLvl int) string {
	var result string
	for _, block := range h.Blocks {
		result += block.Render(indentLvl)
	}
	return result
}

func (h *HorizontalBlock) Type() BlockType {
	return h.BlockType
}

type VerticalBlock struct {
	Blocks    []IBlock
	BlockType BlockType
}

func (v *VerticalBlock) Render(indentLvl int) string {
	var buf bytes.Buffer
	for i, block := range v.Blocks {
		if i != 0 {
			buf.WriteString("\n")
			buf.WriteString(strings.Repeat("\t", indentLvl))
		}
		buf.WriteString(block.Render(indentLvl))
		buf.WriteString("\n")
	}
	return buf.String()
}

func (v *VerticalBlock) Type() BlockType {
	return v.BlockType
}

type IndentBlock struct {
	Block IBlock
}

func (i *IndentBlock) Render(indentLvl int) string {
	return i.Block.Render(indentLvl + 1)
}

func (i *IndentBlock) Type() BlockType {
	return NULL
}

type PrimitiveBlock struct {
	Content   string
	BlockType BlockType
}

func (p *PrimitiveBlock) Render(indentLvl int) string {
	return p.Content
}

func (p *PrimitiveBlock) Type() BlockType {
	return p.BlockType
}
