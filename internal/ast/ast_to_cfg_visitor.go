package ast

import (
	"fmt"
	"strings"

	"github.com/28Pollux28/log6302-parser/utils"
)

// VisitorCFG permet de construire un Control Flow Graph (CFG) à partir d'un AST
type VisitorCFG struct {
	CFG           *CFG
	LastNode      *CFGNode     // Garde une référence au dernier nœud visité
	NodeStack     *utils.Stack // Garde une référence aux structures de contrôle
	NextLabel     string       // Label pour les edges
	currFuncScope int
	FuncStack     *utils.Stack
	ExitNode      *CFGNode
}

// NewVisitorCFG initialise un nouveau visiteur CFG
func NewVisitorCFG() *VisitorCFG {
	v := &VisitorCFG{
		CFG:           NewCFG(),
		NodeStack:     utils.NewStack(),
		NextLabel:     "",
		currFuncScope: 0,
		FuncStack:     utils.NewStack(),
		ExitNode:      nil,
	}
	v.CFG.AddNode("main", "Entry")
	v.LastNode = v.CFG.Nodes[0]
	return v
}

// VisitNode construit les relations entre les nœuds pour créer le CFG
func (v *VisitorCFG) VisitNode(n *Node) {
	if n == nil {
		return
	}

	// switch sur le type de nœud
	switch n.Kind {
	case "binary_expression", "assignment_expression":
		visitBinOP(v, n)
	case "name", "\u003c", "\u003e", "integer", "string_content", "=", "text", "boolean", "variable":
		visitSimple(v, n)
	case "if_statement":
		visitIf(v, n)
	case "while_statement":
		visitWhile(v, n)
	case "break_statement":
		visitBreak(v, n)
	case "continue_statement":
		visitContinue(v, n)
	case "return":
		visitReturn(v, n)
	case "parenthesized_expression":
		visitParenthesizedExpression(v, n)
	case "else_clause", "compound_statement", "program", "variable_name", "expression_statement", "string", "argument", "return_statement":
		visitChilden(v, n)
	case "echo_statement":
		visitStatement(v, n)
	case "encapsed_string":
		visitEncapsedString(v, n)
	case "function_call_expression":
		visitFuncCall(v, n)
	case "function_definition":
		visitFuncDef(v, n)
	case "arguments":
		visitArgs(v, n)
	case "exit_statement":
		visitExit(v, n)
	case "}", "{", "\"", ";", "$", "php_tag", "'", "(", ")", ",", ":", "function", "formal_parameters", "else", "comment":
		return
	default:
		fmt.Printf("Node kind not implemented: %s\n", n.Kind)
		visitChilden(v, n)
	}
}

// BuildCFGFromAST construit un CFG à partir d'une racine AST
func BuildCFGFromAST(root *Node) *CFG {
	visitor := NewVisitorCFG()
	root.accept(visitor)
	return visitor.CFG
}

func visitBinOP(v *VisitorCFG, n *Node) {

	// Visite des enfants
	n.Descendants[0].accept(v)
	n.Descendants[2].accept(v)

	// Création d'un nœud pour l'opération binaire
	node := v.CFG.AddNode(sanitizeHTML(n.Descendants[1].Text), renameKinds(n.Kind))

	// Ajout d'un edge entre le dernier nœud visité et le nœud de l'opération binaire
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

}

func visitSimple(v *VisitorCFG, n *Node) {
	// Création d'un nœud pour la variable
	node := v.CFG.AddNode(sanitizeHTML(n.Text), renameKinds(n.Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""
}

func visitIf(v *VisitorCFG, n *Node) {
	// Création d'un nœud pour le if
	node := v.CFG.AddNode(renameKinds(n.Descendants[0].Kind))

	// Ajout d'un edge entre le dernier nœud visité et le nœud du if
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

	// Visite de la parenthesized_expression
	n.Descendants[1].accept(v)

	// Ajout du nœud de la condition
	conditionNode := v.CFG.AddNode("Condition")
	v.CFG.AddEdge(v.LastNode, conditionNode, v.NextLabel)
	v.LastNode = conditionNode
	v.NextLabel = "True"

	// Ajout d'un nœud pour la fin du if
	endNode := v.CFG.AddNode("IfEnd")
	endNode.SetColor("red")
	endNode.IsSpread = true

	// Visite du bloc then
	n.Descendants[2].accept(v)

	// Ajout d'un edge entre le dernier nœud visité et le nœud de fin du if
	if v.LastNode != nil && v.LastNode.ID > endNode.ID {
		v.CFG.AddEdge(v.LastNode, endNode, "")
	}
	// Visite du bloc else si présent
	if len(n.Descendants) == 4 {
		v.LastNode = conditionNode
		v.NextLabel = "False"
		n.Descendants[3].accept(v)
		v.CFG.AddEdge(v.LastNode, endNode, "")
	} else {
		v.CFG.AddEdge(conditionNode, endNode, "False")
	}

	v.LastNode = endNode

}

func visitWhile(v *VisitorCFG, n *Node) {
	// Création d'un nœud pour le while
	node := v.CFG.AddNode(renameKinds(n.Descendants[0].Kind))

	// Ajout d'un edge entre le dernier nœud visité et le nœud du while
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

	// Visite de la parenthesized_expression
	n.Descendants[1].accept(v)

	// Ajout du nœud de la condition
	conditionNode := v.CFG.AddNode("Condition")
	v.CFG.AddEdge(v.LastNode, conditionNode, v.NextLabel)
	v.LastNode = conditionNode
	v.NextLabel = "True"

	// Ajout d'un nœud pour la fin du while
	endNode := v.CFG.AddNode("WhileEnd")
	endNode.SetColor("red")
	endNode.IsSpread = true
	v.CFG.AddEdge(conditionNode, endNode, "False")

	// sauvegarde du nœud de la condition
	v.NodeStack.Push(utils.Pair{First: node, Second: endNode})

	// Visite du bloc then
	n.Descendants[2].accept(v)

	// Ajout d'un edge entre le dernier nœud visité et le nœud du while
	v.CFG.AddEdge(v.LastNode, node, "")

	v.NodeStack.Pop()
	v.LastNode = endNode
}

func visitBreak(v *VisitorCFG, n *Node) {
	// Création d'un nœud pour le break
	node := v.CFG.AddNode(renameKinds(n.Descendants[0].Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

	// Récupération du nœud de fin du while
	if !v.NodeStack.IsEmpty() {
		pair := v.NodeStack.Peek().(utils.Pair)
		endNode := pair.Second.(*CFGNode)
		// Ajout d'un edge entre le dernier nœud visité et le nœud de fin du while
		v.CFG.AddEdge(v.LastNode, endNode, "")
	}

	v.LastNode = nil
}

func visitContinue(v *VisitorCFG, n *Node) {
	// Création d'un nœud pour le continue
	node := v.CFG.AddNode(renameKinds(n.Descendants[0].Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

	// Récupération du nœud de condition du while
	pair := v.NodeStack.Peek().(utils.Pair)
	whileNode := pair.First.(*CFGNode)

	// Ajout d'un edge entre le dernier nœud visité et le nœud de condition du while
	v.CFG.AddEdge(v.LastNode, whileNode, "")
	v.LastNode = nil
}

func visitReturn(v *VisitorCFG, n *Node) {
	// Création d'un nœud pour le return
	node := v.CFG.AddNode(renameKinds(n.Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

	// todo : Visite de l'expression de retour

	// Exit the function
	exitNode := v.FuncStack.Peek().(*CFGNode)
	v.CFG.AddEdge(v.LastNode, exitNode, v.NextLabel)
	v.LastNode = nil
	v.NextLabel = ""
}

func visitParenthesizedExpression(v *VisitorCFG, n *Node) {
	n.Descendants[1].accept(v)
}

func visitChilden(v *VisitorCFG, n *Node) {
	// Visite des enfants du bloc
	for _, child := range n.Descendants {
		child.accept(v)
	}
	if n.Kind == "program" {
		// Ajout d'un nœud de sortie
		exitNode := v.CFG.AddNode("Exit")
		exitNode.SetColor("red")
		exitNode.IsSpread = true
		v.CFG.AddEdge(v.LastNode, exitNode, "")
		v.LastNode = exitNode
	}
}

func visitStatement(v *VisitorCFG, n *Node) {
	// Création d'un nœud pour le statement
	node := v.CFG.AddNode(renameKinds(n.Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

	// Node pour la liste des arguments
	argsNode := v.CFG.AddNode("ArgumentList")
	v.CFG.AddEdge(node, argsNode, "")
	v.LastNode = argsNode
	v.NextLabel = ""

	// Visite des arguments (tous les enfants sauf le premier)
	for _, child := range n.Descendants[1:] {
		child.accept(v)
	}

	// Propagation valeur de retour
	// retNode := v.CFG.AddNode("Argument")
	// retNode.SetColor("red")
	// retNode.IsSpread = true
	// v.CFG.AddEdge(v.LastNode, retNode, "")
	// v.LastNode = retNode
	// v.NextLabel = ""
}

func visitEncapsedString(v *VisitorCFG, n *Node) {
	n.Descendants[1].accept(v)
}

func visitFuncCall(v *VisitorCFG, n *Node) {
	name := n.Descendants[0].Text
	// Création d'un nœud pour la fonction
	callNode := v.CFG.AddNode(name, "FunctionCall")

	callBeginNode := v.CFG.AddNode(name, "CallBegin")
	callBeginNode.SetColor("red")
	callBeginNode.IsSpread = true

	callEndNode := v.CFG.AddNode(name, "CallEnd")
	callEndNode.SetColor("red")
	callEndNode.IsSpread = true

	idNode := v.CFG.AddNode(name, "Id")
	v.CFG.AddEdge(v.LastNode, callNode, v.NextLabel)
	v.CFG.AddEdge(callNode, idNode, "")
	v.LastNode = idNode
	v.NextLabel = ""

	// Visite des arguments
	n.Descendants[1].accept(v)

	// Propagation valeur de retour
	// argNode := v.CFG.AddNode("Argument")
	// argNode.SetColor("red")
	// argNode.IsSpread = true
	// v.CFG.AddEdge(v.LastNode, argNode, v.NextLabel)
	// v.NextLabel = ""

	retNode := v.CFG.AddNode("RetValue")
	retNode.SetColor("red")
	retNode.IsSpread = true

	v.CFG.AddEdge(v.LastNode, callBeginNode, v.NextLabel)
	v.CFG.AddEdge(callBeginNode, callEndNode, "")
	v.CFG.AddEdge(callEndNode, retNode, "")
	v.LastNode = retNode
	v.NextLabel = ""
}

func visitArgs(v *VisitorCFG, n *Node) {
	node := v.CFG.AddNode("ArgumentList")
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""
	visitChilden(v, n)
}

func visitArg(v *VisitorCFG, n *Node) {
	visitChilden(v, n)

	// Argument node
	argNode := v.CFG.AddNode("Argument")
	argNode.SetColor("red")
	argNode.IsSpread = true
	v.CFG.AddEdge(v.LastNode, argNode, v.NextLabel)
	v.LastNode = argNode
	v.NextLabel = ""
}

func visitFuncDef(v *VisitorCFG, n *Node) {
	lastNode := v.LastNode

	// Create entry node
	node := v.CFG.AddNode(n.Descendants[1].Text, "Entry")
	v.CFG.AddEdge(nil, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""
	node.SetDead(false)

	// add function scope
	v.currFuncScope++
	funcScopeNode := v.CFG.AddNode(n.Descendants[1].Text, "FunctionStatement")
	v.CFG.SetFuncScope(funcScopeNode, v.currFuncScope)
	v.CFG.AddEdge(node, funcScopeNode, "")
	v.LastNode = funcScopeNode

	// add exit node
	exitNode := v.CFG.AddNode("Exit")
	exitNode.SetColor("red")
	exitNode.IsSpread = true

	v.FuncStack.Push(exitNode)

	// visit children
	visitChilden(v, n)

	// add edge to exit node
	v.CFG.AddEdge(v.LastNode, exitNode, v.NextLabel)
	v.NextLabel = ""

	// remove function scope
	v.FuncStack.Pop()

	v.LastNode = lastNode
}

func visitExit(v *VisitorCFG, n *Node) {
	// Check if the exit node already exists
	if v.ExitNode == nil {
		v.ExitNode = v.CFG.AddNode("Dead")
		v.ExitNode.SetColor("red")
		v.ExitNode.IsSpread = true
	}

	// Function call node
	node := v.CFG.AddNode("exit", "FunctionCall")
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.NextLabel = ""

	// Id node
	idNode := v.CFG.AddNode("exit", "Id")
	v.CFG.AddEdge(node, idNode, "")

	// Argument list node
	argsNode := v.CFG.AddNode("ArgumentList")
	v.CFG.AddEdge(idNode, argsNode, "")

	// Argument value node
	argValNode := v.CFG.AddNode(n.Descendants[2].Text, n.Descendants[2].Kind)
	v.CFG.AddEdge(argsNode, argValNode, "")

	// Argument node
	argNode := v.CFG.AddNode("Argument")
	argNode.SetColor("red")
	argNode.IsSpread = true
	v.CFG.AddEdge(argValNode, argNode, "")

	// Add edge to exit node
	v.CFG.AddEdge(argNode, v.ExitNode, v.NextLabel)
	v.LastNode = nil
	v.NextLabel = ""
}

func sanitizeHTML(input string) string {
	// Créer un map pour les substitutions
	replacements := map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"\"": "&quot;",
		"'":  "&apos;",
	}

	// Remplacer les caractères spéciaux
	for old, new := range replacements {
		input = strings.ReplaceAll(input, old, new)
	}

	return input
}

func renameKinds(kind string) string {
	switch kind {
	case "integer":
		return "IntegerLiteral"
	case "string_content":
		return "StringExpression"
	case "variable_name":
		return "Variable"
	case "name":
		return "Id"
	case "if":
		return "If"
	case "while":
		return "While"
	case "echo_statement":
		return "EchoStatement"
	case "binary_expression":
		return "RelOp"
	case "assignment_expression":
		return "BinOp"
	case "break":
		return "Break"
	case "continue":
		return "Continue"
	case "text":
		return "Html"
	default:
		return kind
	}
}
