package day5

type CargoStack struct {
	height int
	cargo  []byte
}

const emptySpace byte = '0'

func CreateCargoStack(cargoList []byte) CargoStack {
	stack := CargoStack{
		height: (len(cargoList)),
		cargo:  cargoList,
	}
	return stack
}

func GetCargoStackInput(cargoDescription string) [][]byte {
	var val [][]byte
	return val
}

func GetCargoValue(cargoLine string, index int) byte {
	// initialise to be empty
	foundCargo := emptySpace

	return foundCargo
}
