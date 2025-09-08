package main

import (
	"fmt"
	"log"

	"github.com/justinhjy1004/go-autodiff/datastructs"
)

func main() {

	n1 := datastructs.Input(6)
	n2 := datastructs.Input(4)

	Logx1, err := datastructs.Log(&n1)

	if err != nil {
		log.Fatal(err)
	}

	datastructs.Sub(datastructs.Add(Logx1, datastructs.Mul(&n1, &n2)), datastructs.Sin(&n2))

	sortedNodes := datastructs.BaseTopologicalSort(&n1)

	//fmt.Println(sortedNodes)

	for _, n := range sortedNodes {
		//fmt.Println(*n)
		fmt.Printf("%s, %f \n", n.Operation, n.Value)
		fmt.Println(n.Grad_wrt_parents)
	}

	//datastructs.PrintTopologicalSort(*v1)
}
