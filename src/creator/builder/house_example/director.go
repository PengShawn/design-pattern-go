package main

type Director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
