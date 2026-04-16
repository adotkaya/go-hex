package ports

//"context"

type GRPCPort interface {
	Run()
	GetAddition()
	GetSubstraction()
	GetMultiplication()
	GetDivision()
}
