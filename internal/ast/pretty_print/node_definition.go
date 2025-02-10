package pretty_print

type Node interface {
	GetChildrenNumber() int
	GetKind() string
	GetText() string
}
