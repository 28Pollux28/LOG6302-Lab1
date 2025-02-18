package ast

import (
	"fmt"
	"strings"
)

type CFGNode struct {
	ID        int
	Label     string
	Title     string
	Children  []*CFGEdge
	Parents   []*CFGEdge
	Color     string
	BgColor   string
	IsDead    bool
	IsSpread  bool
	FuncScope int
}

type CFGEdge struct {
	From *CFGNode
	To   *CFGNode
	Type string
}

type CFG struct {
	Nodes []*CFGNode
	Edges []*CFGEdge
}

func NewCFG() *CFG {
	return &CFG{
		Nodes: []*CFGNode{},
		Edges: []*CFGEdge{},
	}
}

func (node *CFGNode) SetColor(color string) {
	node.Color = color
}

func (node *CFGNode) SetBgColor(color string) {
	node.BgColor = color
}

func (node *CFGNode) SetDead(isDead ...bool) {
	node.IsDead = len(isDead) == 0 || isDead[0]
	if node.IsDead {
		node.SetColor("gray")
		node.SetBgColor("lightgray")
	} else {
		if node.IsSpread {
			node.SetColor("red")
		} else {
			node.SetColor("black")
		}
		node.SetBgColor("white")
	}
}

func (cfg *CFG) AddNode(label string, title ...string) *CFGNode {
	node := &CFGNode{
		ID:        len(cfg.Nodes) + 1,
		Label:     label,
		Color:     "black",
		BgColor:   "white",
		FuncScope: -1,
	}

	if len(title) > 0 {
		node.Title = title[0]
	}

	cfg.Nodes = append(cfg.Nodes, node)
	return node
}

func (cfg *CFG) SetFuncScope(node *CFGNode, scope int) {
	node.FuncScope = scope
}

func (cfg *CFG) AddEdge(from, to *CFGNode, edgeType string) {
	if from == nil {
		to.SetDead(true)
		return
	}

	// if all parents are dead, then the node is dead
	if from.IsDead {
		isDead := true
		for _, parent := range to.Parents {
			if !parent.From.IsDead {
				isDead = false
				break
			}
		}
		to.SetDead(isDead)
	} else {
		to.SetDead(false)
	}
	edge := &CFGEdge{
		From: from,
		To:   to,
		Type: edgeType,
	}
	from.Children = append(from.Children, edge)
	to.Parents = append(to.Parents, edge)
	cfg.Edges = append(cfg.Edges, edge)
}

// GenerateDOT génère une représentation Graphviz du CFG sous forme de chaîne
func (cfg *CFG) GenerateDOT() string {
	var builder strings.Builder

	// Début du graphe
	builder.WriteString("digraph CFG {\n")
	builder.WriteString("  node [shape=none style=\"rounded,filled\" fillcolor=white fontname=\"Arial\" fontsize=12];\n")

	// Ajout des nœuds
	for _, node := range cfg.Nodes {

		if node.Title != "" {
			builder.WriteString(fmt.Sprintf(
				"  %d [label=<<TABLE border='1' color='%s' BGCOLOR='%s' cellspacing='0' cellpadding='10' style='rounded'>"+
					"<TR><TD border='0'>%d</TD><TD border='0'><B>%s</B></TD></TR>"+
					"<HR/><TR><TD border='0' colspan='2'>%s</TD></TR></TABLE>>];\n", node.ID, node.Color, node.BgColor, node.ID, node.Title, node.Label))
		} else {
			builder.WriteString(fmt.Sprintf("  %d [shape=box label=<%d    <B>%s</B>> color=%s fillcolor=%s];\n", node.ID, node.ID, node.Label, node.Color, node.BgColor))
		}
	}

	// Ajout des arêtes
	for _, edge := range cfg.Edges {
		edgeLabel := ""
		if edge.Type != "" {
			edgeLabel = fmt.Sprintf(" [label=\"%s\"]", edge.Type)
		}
		builder.WriteString(fmt.Sprintf("  %d -> %d%s [weight=10];\n", edge.From.ID, edge.To.ID, edgeLabel))
	}

	// Fin du graphe
	builder.WriteString("}\n")

	return builder.String()
}
