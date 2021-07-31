package builder

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) setNumFoor() {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		foor:       b.floor,
	}
}
