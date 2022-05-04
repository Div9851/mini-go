package parser

type NodeKind int

const (
	ND_NUM NodeKind = iota
	ND_ADD
	ND_SUB
	ND_MUL
	ND_DIV
	ND_UNARY_MINUS
)

type Node struct {
	Kind NodeKind
	Num  int
	Lhs  *Node
	Rhs  *Node
}

func newNode(kind NodeKind) *Node {
	node := new(Node)
	node.Kind = kind
	return node
}
