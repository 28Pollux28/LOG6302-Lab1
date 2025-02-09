package test

import (
	"bytes"
	"strings"
)

type IBlock interface {
	Render(indentLvl int) string
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

type VerticalBlock struct {
	Blocks    []IBlock
	BlockType BlockType
}

func (v *VerticalBlock) Render(indentLvl int) string {
	var buf bytes.Buffer
	for _, block := range v.Blocks {
		buf.WriteString(strings.Repeat("\t", indentLvl)) // Indentation
		buf.WriteString(block.Render(indentLvl))
		buf.WriteString("\n")
	}
	return buf.String()
}

type IndentBlock struct {
	Block IBlock
}

func (i *IndentBlock) Render(indentLvl int) string {
	return i.Block.Render(indentLvl + 1)
}

type PrimitiveBlock struct {
	Content   string
	BlockType BlockType
}

func (p *PrimitiveBlock) Render(indentLvl int) string {
	return p.Content
}
