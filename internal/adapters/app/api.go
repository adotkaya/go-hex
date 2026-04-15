package api

import (
	"fmt"
	"go-hex/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPort
}

func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (apiar *Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apiar.arith.Addition(a, b)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return answer, nil
}

func (apiar *Adapter) GetSubstraction(a, b int32) (int32, error) {
	answer, err := apiar.arith.Substraction(a, b)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return answer, nil
}

func (apiar *Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apiar.arith.Multiplication(a, b)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return answer, nil
}

func (apiar *Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apiar.arith.Division(a, b)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return answer, nil
}
