package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}
	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
