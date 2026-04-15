package main

import (
	"fmt"
	"go-hex/core/arithmetic"
)

func main() {
	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}
