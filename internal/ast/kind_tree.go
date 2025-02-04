package ast

import (
	"regexp"
)

type KindTreeAttributes struct {
	Text      *string `json:"text"`
	TextRegex *string `json:"text_regex"`
}

type KindTree struct {
	Name       string              `json:"name"`
	Kind       string              `json:"kind"`
	Attributes *KindTreeAttributes `json:"attributes"`
	Children   []*KindTree         `json:"children"`
}

func NewKindTree(kind string, attributes *KindTreeAttributes) *KindTree {
	return &KindTree{
		Kind:       kind,
		Attributes: attributes,
		Children:   []*KindTree{},
	}
}

func (kt *KindTree) AddChild(kind string, attributes *KindTreeAttributes) {
	kt.Children = append(kt.Children, NewKindTree(kind, attributes))
}

func (kt *KindTree) AddChildTree(child *KindTree) {
	kt.Children = append(kt.Children, child)
}

func (kt *KindTree) Match(n *Node) bool {
	if n.Attributes != nil && n.Attributes[kt.Name+"matched"].V == true {
		return false
	}
	if kt.Kind != "any" && n.Kind != kt.Kind {
		return false
	}
	if kt.Attributes != nil {
		if !kt.Attributes.Match(n) {
			return false
		}
	}
	for _, child := range kt.Children {
		found := false
		for _, descendant := range n.Descendants {
			if child.Match(descendant) {
				found = true
				descendant.Attributes[kt.Name+"matched"] = Attribute[any]{V: true}
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (kta *KindTreeAttributes) Match(n *Node) bool {
	if kta.Text != nil {
		if n.Text != *kta.Text {
			return false
		}
	}
	if kta.TextRegex != nil {
		if !regexp.MustCompile(*kta.TextRegex).MatchString(n.Text) {
			return false
		}
	}
	return true
}
