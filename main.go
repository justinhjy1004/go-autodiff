package main

import (
	"fmt"
	"log"

	"github.com/justinhjy1004/go-autodiff/datastructs"
)

func main() {

	n1 := datastructs.Input(3)
	n2 := datastructs.Input(4)

	Logx1, err := datastructs.Log(&n1)

	if err != nil {
		log.Fatal(err)
	}

	/* y = sin(ln(n1) + (n1 * n2) - sin(n2))
	   dy/dn1 = ln(n1) + (n1 * n2) - sin(n2)
	   dy/dn2 = n1 - cos(n2) */
	datastructs.Sub(datastructs.Add(Logx1, datastructs.Mul(&n1, &n2)), datastructs.Sin(&n2))

	//datastructs.Add(datastructs.Sub(&n1, &n2), Logx1)

	sortedNodes := datastructs.BaseTopologicalSort(&n1)

	var ans float64 = datastructs.InductionDerivative(sortedNodes[0], sortedNodes)

	fmt.Println(ans)

}
