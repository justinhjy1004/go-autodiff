package main

import (
	"fmt"

	"github.com/justinhjy1004/go-autodiff/datastructs"
)

func main() {

	n1 := datastructs.Input(6)
	n2 := datastructs.Input(4)
	n3 := datastructs.Input(0.5)

	datastructs.Add(datastructs.Add(datastructs.Mul(&n1, &n2), datastructs.Sin(&n3)), datastructs.Cos(&n1))

	sortedNodes := datastructs.BaseTopologicalSort(&n1)

	//fmt.Println(sortedNodes)

	for _, n := range sortedNodes {
		fmt.Println(*n)
		fmt.Printf("%s, %f \n", n.Operation, n.Value)
	}

	//datastructs.PrintTopologicalSort(*v1)
}
