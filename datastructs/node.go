package datastructs

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

type Node struct {
	Value              float64
	Parents            []Node
	Operation          string
	Grad_wrt_parents   []float64
	Partial_derivative float64
}

func (n Node) GetParents() []Node {
	return n.Parents
}

func Input(x float64) Node {

	input_node := Node{Value: x, Operation: "input", Partial_derivative: 0}

	return input_node
}

func Mul(x Node, y Node) Node {

	// Only fo constants though
	value := x.Value * y.Value
	operation := "mul"
	grad_wrt_parents := []float64{y.Value, x.Value}
	parents := []Node{x, y}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents, Partial_derivative: 0}

	return result

}

func Add(x Node, y Node) Node {

	value := x.Value + y.Value
	operation := "add"
	grad_wrt_parents := []float64{1, 1}
	parents := []Node{x, y}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents, Partial_derivative: 0}

	return result
}

func Sub(x Node, y Node) Node {

	value := x.Value - y.Value
	operation := "sub"
	grad_wrt_parents := []float64{1, -1}
	parents := []Node{x, y}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents, Partial_derivative: 0}

	return result
}

// TODO
func Log(x Node) (Node, error) {

	if x.Value <= 0 {
		return Node{}, errors.New("Input of log is <= 0")
	}

	value := math.Log(x.Value)
	operation := "log"
	grad_wrt_parents := []float64{1 / x.Value}
	parents := []Node{x}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	return result, nil

}

func Sin(x Node) Node {

	value := math.Sin(x.Value)
	operation := "sin"
	grad_wrt_parents := []float64{math.Cos(x.Value)}
	parents := []Node{x}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	return result
}

func Cos(x Node) Node {

	value := math.Cos(x.Value)
	operation := "cos"
	grad_wrt_parents := []float64{-math.Sin(x.Value)}
	parents := []Node{x}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	return result
}

func inductiveTopologicalSort(n Node, result *[]*Node) {

	if len(n.Parents) == 0 {

		if !slices.Contains(*result, &n) {
			*result = append(*result, &n)
		}

		return

	}

	for _, parent := range n.Parents {

		if !slices.Contains(*result, &parent) {
			inductiveTopologicalSort(parent, result)
		}
	}

	*result = append([]*Node{&n}, *result...)

}

// Given a single output, backward trace the dependencies
func BaseTopologicalSort(n Node) []*Node {

	result := []*Node{}

	for _, parent := range n.Parents {

		if !slices.Contains(result, &parent) {
			inductiveTopologicalSort(parent, &result)
		}

	}

	result = append([]*Node{&n}, result...)

	return result
}

// Print Topological Sort Outcome
func PrintTopologicalSort(node Node) {

	for _, n := range BaseTopologicalSort(node) {
		fmt.Printf("%s, %f \n", n.Operation, n.Value)
	}

}
