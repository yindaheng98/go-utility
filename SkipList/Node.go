package SkipList

//The node consists of SkipList.
type Node struct {
	data float64
	prev []*Node
	next []*Node
}

//Get the value of the node.
func (n Node) GetData() float64 {
	return n.data
}

func newNode(Data float64, level uint64) *Node {
	if level < 1 {
		level = 1
	}
	return &Node{Data, make([]*Node, level), make([]*Node, level)}
}
