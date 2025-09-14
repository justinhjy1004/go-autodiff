package datastructs

import (
	"errors"
	"math"
	"slices"
)

// Node struct to implement computation graph
type Node struct {
	Value            float64 //holds value of current computational step
	Parents          []*Node // parents of the current node
	Children         []*Node
	Operation        string    // what operation was performed on the value of the parents
	Grad_wrt_parents []float64 // dnode/dparents
}

// Input node
func Input(x float64) Node {

	input_node := Node{Value: x, Operation: "input"}

	return input_node
}

// Multiply node, binary operation that leads to a new node
// WARNING: I am assuming that x and y are not f(x) and f(y) but constants wrt to the derived input
func Mul(x *Node, y *Node) *Node {

	// Only fo constants though
	value := x.Value * y.Value
	operation := "mul"
	grad_wrt_parents := []float64{y.Value, x.Value}
	parents := []*Node{x, y}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	x.Children = append(x.Children, &result)
	y.Children = append(y.Children, &result)

	return &result

}

// Addition node, also a binary operation
func Add(x *Node, y *Node) *Node {

	value := x.Value + y.Value
	operation := "add"
	grad_wrt_parents := []float64{1, 1}
	parents := []*Node{x, y}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	x.Children = append(x.Children, &result)
	y.Children = append(y.Children, &result)

	return &result
}

// Subtraction node, also a binary operation
func Sub(x *Node, y *Node) *Node {

	value := x.Value - y.Value
	operation := "sub"
	grad_wrt_parents := []float64{1, -1}
	parents := []*Node{x, y}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	x.Children = append(x.Children, &result)
	y.Children = append(y.Children, &result)

	return &result
}

// Calculates the natural log
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

// Sin function
func Sin(x *Node) *Node {

	value := math.Sin(x.Value)
	operation := "sin"
	grad_wrt_parents := []float64{math.Cos(x.Value)}
	parents := []*Node{x}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	x.Children = append(x.Children, &result)

	return &result
}

// Cosine function
func Cos(x *Node) *Node {

	value := math.Cos(x.Value)
	operation := "cos"
	grad_wrt_parents := []float64{-math.Sin(x.Value)}
	parents := []*Node{x}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	x.Children = append(x.Children, &result)

	return &result
}

// Polynomial term
func Pol(x *Node, power float64, coefficient float64) *Node {

	value := coefficient * math.Pow(x.Value, power)
	operation := "pol"
	grad_wrt_parents := []float64{coefficient * power * math.Pow(x.Value, power-1)}
	parents := []*Node{x}

	result := Node{Value: value, Operation: operation, Grad_wrt_parents: grad_wrt_parents, Parents: parents}

	x.Children = append(x.Children, &result)

	return &result
}

/*
	Inductive step of topological sort (see function below)
*/

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

/*
Input: Pointer of the input of equation. (ie. Creates the topological sort of the computation graph wrt dy/dinput)
Returns: Backward trace the dependencies starting from output
*/
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

/*
Input: Takes the first node of the inverted topological sort (output), and the topological sorted nodes
Output: The partial derivative wrt to the input
*/
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

/*
Input: Takes a slice of Node pointers. They should be inputs Nodes
Output: slice of partial derivatives in the same order
*/
func Partial_Derivative(inputs []*Node) ([]float64, error) {

	partial_derivatives := make([]float64, len(inputs))

	for i, input := range inputs {

		if input.Operation != "input" {
			return partial_derivatives, errors.New("what are you even doing bro? we need inputs")
		}

		topSort := BaseTopologicalSort(input)
		partial_derivatives[i] = InductionDerivative(topSort[0], topSort)

	}

	return partial_derivatives, nil
}
