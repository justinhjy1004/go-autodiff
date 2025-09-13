package datastructs

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

type Node struct {
	Value              float64
	Parents            []*Node
	Children           []*Node
	Operation          string
	Grad_wrt_parents   []float64
	Partial_derivative float64
}

func (n Node) GetParents() []*Node {
	return n.Parents
}

func Input(x float64) Node {

	input_node := Node{Value: x, Operation: "input", Partial_derivative: 0}

	return input_node
}

func Mul(x *Node, y *Node) *Node {

	// Only fo constants though
	value := x.Value * y.Value
	operation := "mul"
	grad_wrt_parents := []float64{y.Value, x.Value}
	parents := []*Node{x, y}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents, Partial_derivative: 0}

	x.Children = append(x.Children, &result)
	y.Children = append(y.Children, &result)

	return &result

}

func Add(x *Node, y *Node) *Node {

	value := x.Value + y.Value
	operation := "add"
	grad_wrt_parents := []float64{1, 1}
	parents := []*Node{x, y}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents, Partial_derivative: 0}

	x.Children = append(x.Children, &result)
	y.Children = append(y.Children, &result)

	return &result
}

func Sub(x *Node, y *Node) *Node {

	value := x.Value - y.Value
	operation := "sub"
	grad_wrt_parents := []float64{1, -1}
	parents := []*Node{x, y}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents, Partial_derivative: 0}

	x.Children = append(x.Children, &result)
	y.Children = append(y.Children, &result)

	return &result
}

// TODO
func Log(x *Node) (*Node, error) {

	if x.Value <= 0 {
		return nil, errors.New("Input of log is <= 0")
	}

	value := math.Log(x.Value)
	operation := "log"
	grad_wrt_parents := []float64{1 / x.Value}
	parents := []*Node{x}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	x.Children = append(x.Children, &result)

	return &result, nil

}

func Sin(x *Node) *Node {

	value := math.Sin(x.Value)
	operation := "sin"
	grad_wrt_parents := []float64{math.Cos(x.Value)}
	parents := []*Node{x}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	x.Children = append(x.Children, &result)

	return &result
}

func Cos(x *Node) *Node {

	value := math.Cos(x.Value)
	operation := "cos"
	grad_wrt_parents := []float64{-math.Sin(x.Value)}
	parents := []*Node{x}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	x.Children = append(x.Children, &result)

	return &result
}

func inductiveTopologicalSort(n *Node, result *[]*Node) {

	if len(n.Children) == 0 {

		if !slices.Contains(*result, n) {
			*result = append(*result, n)
		}

		return

	}

	for _, child := range n.Children {

		if !slices.Contains(*result, child) {
			inductiveTopologicalSort(child, result)
		}
	}

	*result = append([]*Node{n}, *result...)

}

// Given a single output, backward trace the dependencies
func BaseTopologicalSort(n *Node) []*Node {

	result := []*Node{}

	for _, child := range n.Children {

		if !slices.Contains(result, child) {
			inductiveTopologicalSort(child, &result)
		}

	}

	result = append([]*Node{n}, result...)

	slices.Reverse(result)

	return result
}

// Print Topological Sort Outcome
func PrintTopologicalSort(node Node) {

	for _, n := range BaseTopologicalSort(&node) {
		fmt.Printf("%s, %f \n", n.Operation, n.Value)
	}

}

func InductionDerivative(currentNode *Node, sortedNodes []*Node) float64 {

	var derivative float64 = 0

	if currentNode.Operation == "input" {

		if slices.Contains(sortedNodes, currentNode) {
			return 1
		} else {
			return 0
		}

	} else {
		for i, parent := range currentNode.Parents {

			if slices.Contains(sortedNodes, parent) {

				derivative = derivative + InductionDerivative(parent, sortedNodes)*currentNode.Grad_wrt_parents[i]

			}
		}

		return derivative
	}

}
