package main

import (
	"fmt"
	"log"

	"github.com/justinhjy1004/go-autodiff/datastructs"
)

func main() {

	fmt.Println("Hello")

	n1 := datastructs.Input(6)
	n2 := datastructs.Input(4)
	n3 := datastructs.Input(0.5)

	LogN2, err := datastructs.Log(n2)

	if err != nil {
		log.Fatal(err)
	}

	v1 := datastructs.Add(LogN2, datastructs.Mul(datastructs.Sin(n3), datastructs.Mul(n1, n2)))

	fmt.Println(datastructs.Base_topological_sort(v1))

	for _, n := range datastructs.Base_topological_sort(v1) {
		fmt.Printf("%s, %f \n", n.Operation, n.Value)
	}
}
