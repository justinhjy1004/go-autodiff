package main

import (
	"fmt"
	"log"

	"github.com/justinhjy1004/go-autodiff/datastructs"
)

func main() {

	// Input Variables
	n1 := datastructs.Input(2)
	n2 := datastructs.Input(5)
	n3 := datastructs.Input(0)

	// Define logs, yeah, I wanted to try error handling and it fucked me with this grossness
	// if someone wants to write a parser for a more "human way of writing math", feel free to do so
	Logx1, err := datastructs.Log(&n1)

	if err != nil {
		log.Fatal(err)
	}

	/* Definind the equation and respective derivatives
	   y = sin(ln(n1) + (n1 * n2) - sin(n2)) - 2(n3)^3
	   dy/dn1 = ln(n1) + (n1 * n2) - sin(n2)
	   dy/dn2 = n1 - cos(n2)
	   dy/dn3 = 6(n3)^2 */
	y := datastructs.Sub(datastructs.Sub(datastructs.Add(Logx1, datastructs.Mul(&n1, &n2)), datastructs.Sin(&n2)), datastructs.Pol(&n3, 3, 2))

	fmt.Printf("Value of y: %f \n", y.Value)

	// Calculate partial derivative of each input
	pd, err := datastructs.Partial_Derivative([]*datastructs.Node{&n1, &n2, &n3})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pd)

}
