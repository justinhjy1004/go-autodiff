package main

import (
	"fmt"

	"github.com/justinhjy1004/go-autodiff/datastructs"
)

func main() {

	fmt.Println("Hello")

	n1 := datastructs.Input(6)
	n2 := datastructs.Input(4)
	n3 := datastructs.Input(0.5)

	v1 := datastructs.Mul(n3, datastructs.Mul(n1, n2))

	fmt.Println(v1)
}
