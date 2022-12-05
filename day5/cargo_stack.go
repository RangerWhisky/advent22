package day5

type CargoStack struct {
	height int
	cargo  []byte
}

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
