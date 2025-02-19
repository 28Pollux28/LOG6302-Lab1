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
	FuncDef       []string
	FuncCall      []string
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
		FuncDef:       []string{},
		FuncCall:      []string{},
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
	case "binary_expression", "assignment_expression", "member_access_expression":
		visitBinOP(v, n)
	case "unary_op_expression":
		visitUnOP(v, n)
	case "if_statement":
		visitIf(v, n)
	case "while_statement":
		visitWhile(v, n)
	case "foreach_statement":
		visitForEach(v, n)
	case "break_statement":
		visitBreak(v, n)
	case "continue_statement":
		visitContinue(v, n)
	case "return_statement":
		visitReturnStatement(v, n)
	case "return":
		visitReturn(v, n)
	case "parenthesized_expression":
		visitParenthesizedExpression(v, n)
	case "echo_statement", "unset_statement", "set_statement":
		visitStatement(v, n)
	case "encapsed_string":
		visitEncapsedString(v, n)
	case "function_call_expression", "member_call_expression":
		visitFuncCall(v, n)
	case "function_definition":
		visitFuncDef(v, n)
	case "arguments":
		visitArgs(v, n)
	case "argument":
		visitArg(v, n)
	case "exit_statement":
		visitExit(v, n)
	case "require_expression", "require_once_expression":
		visitRequire(v, n)
	case "global":
		visitGlobal(v, n)
	case "method_declaration":
		visitMethod(v, n)
	case "subscript_expression":
		visitSubsExpr(v, n)
	case "for_statement":
		visitFor(v, n)
	case "switch_statement":
		visitSwitch(v, n)
	case "case_statement":
		visitCase(v, n)
	case "default_statement":
		visitDefault(v, n)
	case "static_modifier", "visibility_modifier", "++", "name", "new", "\u003c", "\u003e", "integer", "string_content", "=", ".=", "text", "boolean", "variable", "->", "cast_type":
		visitSimple(v, n)
	case "update_expression", "switch_block", "else_clause", "compound_statement", "program", "variable_name", "expression_statement", "string", "text_interpolation", "pair", "global_declaration", "array_element_initializer", "colon_block", "object_creation_expression":
		visitChilden(v, n)
	case "array_creation_expression", "cast_expression":
		visitSelfAndChildren(v, n)
	case "}", "{", "\"", ";", "$", "php_tag", "'", "(", ")", ",", ":", "function", "formal_parameters", "else", "comment", "default", "require", "]", "[", "ERROR", "require_once", "as", "=>", "?>", "array", "set", "unset", "escape_sequence", "switch", "case":
		return
	default:
		//fmt.Printf("Node kind not implemented: %s\n", n.Kind)
		visitChilden(v, n)
	}
}

// BuildCFGFromAST construit un CFG à partir d'une racine AST
func BuildCFGFromAST(root *Node, checkInter ...bool) *CFG {
	visitor := NewVisitorCFG()
	root.accept(visitor)
	if len(checkInter) > 0 && checkInter[0] {
		_ = visitor.CheckDeadCodeInter()
	}

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

	offset := 0
	// Visite du bloc then
	if n.Descendants[2].Kind == "comment" {
		offset = 1

	}
	n.Descendants[2+offset].accept(v)

	// Ajout d'un edge entre le dernier nœud visité et le nœud de fin du if
	if v.LastNode != nil && v.LastNode.ID > endNode.ID {
		v.CFG.AddEdge(v.LastNode, endNode, "")
	}
	// Visite du bloc else si présent
	if len(n.Descendants) == 4+offset {
		v.LastNode = conditionNode
		v.NextLabel = "False"
		n.Descendants[3+offset].accept(v)
		v.CFG.AddEdge(v.LastNode, endNode, "")
	} else {
		v.CFG.AddEdge(conditionNode, endNode, "False")
		if !conditionNode.IsDead {
			endNode.SetDead(false)
		}
	}

	v.LastNode = endNode

}

func visitSwitch(v *VisitorCFG, n *Node) {
	node := v.CFG.AddNode(renameKinds(n.Descendants[0].Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

	// Create end node
	endNode := v.CFG.AddNode("EndSwitch")
	endNode.SetColor("red")
	endNode.IsSpread = true
	v.NodeStack.Push(utils.Pair{First: node, Second: endNode})

	// Visit of parenthesized_expression
	n.Descendants[1].accept(v)

	// add Variable spread node
	varNode := v.CFG.AddNode("#swtich_value_0", "Variable")
	v.CFG.AddEdge(v.LastNode, varNode, v.NextLabel)
	varNode.IsSpread = true
	varNode.SetColor("red")

	// add binary node "=" once
	binNode := v.CFG.AddNode("=", "BinOp")
	v.CFG.AddEdge(varNode, binNode, "")
	binNode.IsSpread = true
	binNode.SetColor("red")

	v.LastNode = binNode
	v.NextLabel = ""

	// visit switch block
	n.Descendants[2].accept(v)

	v.NodeStack.Pop()

	v.LastNode = endNode
}

func visitCase(v *VisitorCFG, n *Node) {
	// Case node
	node := v.CFG.AddNode(renameKinds(n.Kind))

	// Var node
	varNode := v.CFG.AddNode("#swtich_value_1", "Variable")
	v.CFG.AddEdge(v.LastNode, varNode, v.NextLabel)
	v.LastNode = varNode
	v.NextLabel = ""
	varNode.IsSpread = true
	varNode.SetColor("red")

	// value node
	n.Descendants[1].accept(v)

	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	node.SetColor("red")
	node.IsSpread = true

	v.LastNode = node
	v.NextLabel = "True"

	// visit children
	for _, child := range n.Descendants[3:] {
		child.accept(v)
	}

	// continue to the next case
	v.LastNode = node
	v.NextLabel = "False"
}

func visitDefault(v *VisitorCFG, n *Node) {
	node := v.CFG.AddNode(renameKinds(n.Kind))
	node.SetColor("red")
	node.IsSpread = true
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

	// visit children
	for _, child := range n.Descendants[1:] {
		child.accept(v)
	}

	// link to end switch
	pair := v.NodeStack.Peek().(utils.Pair)
	endNode := pair.Second.(*CFGNode)
	v.CFG.AddEdge(v.LastNode, endNode, "")
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

func visitForEach(v *VisitorCFG, n *Node) {
	// Création d'un nœud pour le foreach
	node := v.CFG.AddNode(renameKinds(n.Descendants[0].Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

	// Ajout d'un nœud pour la fin du foreach
	endNode := v.CFG.AddNode("ForeachEnd")
	endNode.SetColor("red")
	endNode.IsSpread = true
	v.CFG.AddEdge(node, endNode, "")

	// empilement du nœud du foreach
	v.NodeStack.Push(utils.Pair{First: node, Second: endNode})

	// Visite de la parenthesized_expression
	n.Descendants[2].accept(v)

	// Ajout d'un edge entre le dernier nœud visité et le nœud du foreach
	v.CFG.AddEdge(v.LastNode, node, "")

	v.NodeStack.Pop()
	v.LastNode = endNode
}

func visitFor(v *VisitorCFG, n *Node) {
	// Création d'un nœud pour le for
	node := v.CFG.AddNode(renameKinds(n.Descendants[0].Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""

	// Ajout d'un nœud pour la fin du for
	endNode := v.CFG.AddNode("ForEnd")
	endNode.SetColor("red")
	endNode.IsSpread = true

	// Visit of assignment_expression
	n.Descendants[2].accept(v)

	// add argument node
	argNode := v.CFG.AddNode("Argument")
	argNode.SetColor("red")
	argNode.IsSpread = true
	v.CFG.AddEdge(v.LastNode, argNode, "")
	v.LastNode = argNode

	// add ForInitEnd node
	initEndNode := v.CFG.AddNode("ForInitEnd")
	initEndNode.SetColor("red")
	initEndNode.IsSpread = true
	v.CFG.AddEdge(v.LastNode, initEndNode, "")
	v.LastNode = initEndNode

	// Visit of binary_expression
	n.Descendants[4].accept(v)

	// Add condition node
	condNode := v.CFG.AddNode("Condition")
	condNode.SetColor("red")
	condNode.IsSpread = true
	v.CFG.AddEdge(v.LastNode, condNode, "")
	v.LastNode = condNode

	// Link to end node
	v.CFG.AddEdge(v.LastNode, endNode, "False")

	// Node increment
	incNode := v.CFG.AddNode("Increment")
	v.LastNode = incNode

	// empilement du nœud du for
	v.NodeStack.Push(utils.Pair{First: incNode, Second: endNode})
	offset := 0
	if n.Descendants[6].Kind == "update_expression" {
		// arg node
		argNode2 := v.CFG.AddNode("Argument")
		argNode2.SetColor("red")
		argNode2.IsSpread = true

		// Visit of update_expression
		n.Descendants[6].accept(v)
		v.CFG.AddEdge(v.LastNode, argNode2, "")
		v.CFG.AddEdge(argNode2, initEndNode, v.NextLabel)
	} else {
		v.CFG.AddEdge(incNode, initEndNode, v.NextLabel)
		offset = -1
	}
	// Visit of compound_statement
	v.LastNode = condNode
	v.NextLabel = "True"
	n.Descendants[8+offset].accept(v)

	v.CFG.AddEdge(v.LastNode, incNode, "")

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
	if !v.NodeStack.IsEmpty() {
		pair := v.NodeStack.Peek().(utils.Pair)
		whileNode := pair.First.(*CFGNode)

		// Ajout d'un edge entre le dernier nœud visité et le nœud de condition du while
		v.CFG.AddEdge(v.LastNode, whileNode, "")
	}
	v.LastNode = nil
}

func visitReturnStatement(v *VisitorCFG, n *Node) {
	// visit all children except the first one
	for _, child := range n.Descendants[1:] {
		child.accept(v)
	}
	// visit the first child (the return keyword)
	n.Descendants[0].accept(v)

}

func visitReturn(v *VisitorCFG, n *Node) {

	node := v.CFG.AddNode(renameKinds(n.Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)

	// Exit the function
	if !v.FuncStack.IsEmpty() {
		exitNode := v.FuncStack.Peek().(*CFGNode)
		v.CFG.AddEdge(v.LastNode, exitNode, v.NextLabel)
	}
	v.LastNode = nil
	v.NextLabel = ""
}

func visitParenthesizedExpression(v *VisitorCFG, n *Node) {
	n.Descendants[1].accept(v)
}

func visitChilden(v *VisitorCFG, n *Node) {
	// Visit children
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

func visitSelfAndChildren(v *VisitorCFG, n *Node) {
	// visit children
	visitChilden(v, n)

	// visit self
	visitSimple(v, n)
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

/*
func visitMemberAcces(v *VisitorCFG, n *Node) {
	if len(n.Descendants) == 3 {
		visitBinOP(v, n)
	} else {
		visitFuncCall(v, n)
	}
}*/

func visitFuncCall(v *VisitorCFG, n *Node) {
	offset := 0
	if n.Kind == "member_call_expression" {
		offset = 2
		n.Descendants[0].accept(v)
	}

	name := n.Descendants[0+offset].Text

	v.FuncCall = append(v.FuncCall, name)
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
	n.Descendants[1+offset].accept(v)

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
	name := n.Descendants[1].Text
	v.FuncDef = append(v.FuncDef, name)

	// Create entry node
	node := v.CFG.AddNode(name, "Entry")
	v.CFG.AddEdge(nil, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""
	node.SetDead(false)

	// add function scope
	v.currFuncScope++
	funcScopeNode := v.CFG.AddNode(name, "FunctionStatement")
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
	v.LastNode = idNode

	if len(n.Descendants) > 2 && n.Descendants[2].Kind != ")" {
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
		v.LastNode = argNode
	}

	// Add edge to exit node
	v.CFG.AddEdge(v.LastNode, v.ExitNode, v.NextLabel)
	v.LastNode = nil
	v.NextLabel = ""
}

func visitRequire(v *VisitorCFG, n *Node) {
	visitChilden(v, n)
	name := n.Descendants[1].Descendants[1].Text
	callBeginNode := v.CFG.AddNode(name, "IncludeBegin")
	callBeginNode.SetColor("red")
	callBeginNode.IsSpread = true

	callEndNode := v.CFG.AddNode(name, "IncludeEnd")
	callEndNode.SetColor("red")
	callEndNode.IsSpread = true

	v.CFG.AddEdge(v.LastNode, callBeginNode, v.NextLabel)
	v.CFG.AddEdge(callBeginNode, callEndNode, "")
	v.LastNode = callEndNode
	v.NextLabel = ""
}

func visitUnOP(v *VisitorCFG, n *Node) {
	// Visite des enfants
	n.Descendants[1].accept(v)

	// Création d'un nœud pour l'opération unaire
	node := v.CFG.AddNode(sanitizeHTML(n.Descendants[0].Text), renameKinds(n.Kind))

	// Ajout d'un edge entre le dernier nœud visité et le nœud de l'opération unaire
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""
}

func visitGlobal(v *VisitorCFG, n *Node) {
	// Création d'un nœud pour le global
	node := v.CFG.AddNode(renameKinds(n.Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""
}

func visitSubsExpr(v *VisitorCFG, n *Node) {
	for _, child := range n.Descendants {
		child.accept(v)
	}

	node := v.CFG.AddNode(renameKinds(n.Kind))
	v.CFG.AddEdge(v.LastNode, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""
}

func visitMethod(v *VisitorCFG, n *Node) {
	// todo finish this
	lastNode := v.LastNode
	// add function scope
	v.currFuncScope++

	name := "Method" + fmt.Sprint(v.currFuncScope)
	v.FuncDef = append(v.FuncDef, name)

	// Create entry node
	node := v.CFG.AddNode(name, "Entry")
	v.CFG.AddEdge(nil, node, v.NextLabel)
	v.LastNode = node
	v.NextLabel = ""
	node.SetDead(false)

	funcScopeNode := v.CFG.AddNode(name, "FunctionStatement")
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

func sanitizeHTML(input string) string {
	// Créer un map pour les substitutions
	replacements := map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"\"": "&quot;",
		"'":  "&apos;",
		"\n": "",
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
	case "assignment_expression", "member_access_expression":
		return "BinOp"
	case "break":
		return "Break"
	case "continue":
		return "Continue"
	case "text":
		return "Html"
	case "unary_op_expression":
		return "UnaryOp"
	case "subscript_expression":
		return "ArrayExpression"
	case "array_creation_expression":
		return "ArrayInitialisation"
	case "cast_expression":
		return "CastExpression"
	case "switch_statement":
		return "Switch"
	case "case_statement":
		return "CaseCondition"
	case "default_statement":
		return "CaseDefault"
	case "for":
		return "For"
	default:
		return kind
	}
}

func (v *VisitorCFG) CheckDeadCodeInter() bool {
	for _, funcDef := range v.FuncDef {
		found := false
		for _, funcCall := range v.FuncCall {
			if funcDef == funcCall {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Function %s is not called\n", funcDef)
			return true
		}
	}
	return false
}
