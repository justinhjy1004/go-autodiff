package datastructs

import "math"

type Node struct {
	Value              float64
	Parents            []Node
	Operation          string
	Grad_wrt_parents   []float64
	Partial_derivative float64
}

func (n Node) Get_parents() []Node {
	return n.Parents
}

func Input(x float64) Node {

	input_node := Node{Value: x, Operation: "input", Partial_derivative: 0}

	return input_node
}

func Mul(x Node, y Node) Node {

	value := x.Value * y.Value
	operation := "mul"
	grad_wrt_parents := []float64{y.Value, x.Value}
	parents := []Node{x, y}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents, Partial_derivative: 0}

	return result

}

// TODO
func Add(x Node, y Node) Node {
	return Node{}
}

// TODO
func Log(x Node) Node {
	return Node{}
}

// TODO
func Sin(x Node) Node {
	return Node{Value: math.Sin(x.Value)}
}
