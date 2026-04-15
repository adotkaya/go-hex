package api

import (
	"go-hex/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPort
}

func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (apiar *Adapter) GetAddition(a, b int32) (int32, error) {
	return apiar.arith.Addition(a, b)
}

func (apiar *Adapter) GetSubstraction(a, b int32) (int32, error) {
	return apiar.arith.Substraction(a, b)
}

func (apiar *Adapter) GetMultiplication(a, b int32) (int32, error) {
	return apiar.arith.Multiplication(a, b)
}

func (apiar *Adapter) GetDivision(a, b int32) (int32, error) {
	return apiar.arith.Division(a, b)
}
