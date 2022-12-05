package day5

type CargoStack struct {
	height int
	cargo  []byte
}

func CreateCargoStack(cargoList []byte) CargoStack {
	stack := CargoStack{
		height: 0,
		cargo:  cargoList,
	}
	return stack
}
