package builder

// Director struct
type Director struct {
	builder IBuilder
}

// Create new Director
func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// Set Builder
func (d *Director) SetBuilder(b *IBuilder) {
	d.builder = b
}

// Build House
func (d *Director) BuildHouse() House {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFoor()
	return d.builder.getHouse()
}
