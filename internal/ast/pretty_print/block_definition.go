package pretty_print

import (
	"bytes"
	"strings"
)

type IBlock interface {
	Render(indentLvl int) string
	Type() BlockType
	GetBlocks() []IBlock
	AppendBlock(block IBlock)
}

type HorizontalBlock struct {
	Blocks    []IBlock
	BlockType BlockType
}

func (h *HorizontalBlock) Render(indentLvl int) string {
	var buf bytes.Buffer
	for _, block := range h.Blocks {
		buf.WriteString(block.Render(indentLvl))
	}
	return buf.String()
}

func (h *HorizontalBlock) Type() BlockType {
	return h.BlockType
}

func (h *HorizontalBlock) GetBlocks() []IBlock {
	return h.Blocks
}

func (h *HorizontalBlock) AppendBlock(block IBlock) {
	h.Blocks = append(h.Blocks, block)
}

type VerticalBlock struct {
	Blocks      []IBlock
	BlockType   BlockType
	IndentFirst bool
}

func (v *VerticalBlock) Render(indentLvl int) string {
	var buf bytes.Buffer
	for i, block := range v.Blocks {
		if i != 0 {
			buf.WriteString("\n")
		}
		if v.IndentFirst || i != 0 {
			buf.WriteString(strings.Repeat("    ", indentLvl))
		}
		buf.WriteString(block.Render(indentLvl))
	}
	return buf.String()
}

func (v *VerticalBlock) Type() BlockType {
	return v.BlockType
}

func (v *VerticalBlock) GetBlocks() []IBlock {
	return v.Blocks
}

func (v *VerticalBlock) AppendBlock(block IBlock) {
	v.Blocks = append(v.Blocks, block)
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

func (i *IndentBlock) GetBlocks() []IBlock {
	return []IBlock{i.Block}
}

func (i *IndentBlock) AppendBlock(_ IBlock) {
	panic("cannot append to IndentBlock")
}

type PrimitiveBlock struct {
	Content   string
	BlockType BlockType
}

func (p *PrimitiveBlock) Render(_ int) string {
	return p.Content
}

func (p *PrimitiveBlock) Type() BlockType {
	return p.BlockType
}

func (p *PrimitiveBlock) GetBlocks() []IBlock {
	return nil
}

func (p *PrimitiveBlock) AppendBlock(_ IBlock) {
	panic("cannot append to PrimitiveBlock")
}
