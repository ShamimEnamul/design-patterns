package main

import "fmt"

// Builder Interface

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "btibuilder" {
		return newBtiBuilder()
	}
	if builderType == "igloobuilder" {
		return newIglooBuilder()
	}

	return nil
}

// concrete builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "wooden window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "wooden door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 2
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// concrete builder

type BtiBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newBtiBuilder() *BtiBuilder {
	return &BtiBuilder{}
}

func (b *BtiBuilder) setWindowType() {
	b.windowType = "wooden window"
}

func (b *BtiBuilder) setDoorType() {
	b.doorType = "wooden door"
}

func (b *BtiBuilder) setNumFloor() {
	b.floor = 2
}

func (b *BtiBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// Product

type House struct {
	windowType string
	doorType   string
	floor      int
}

// director

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setNumFloor()
	d.builder.setDoorType()
	return d.builder.getHouse()
}

// client
func main() {
	builder := getBuilder("btibuilder") //newBtiBuilder()
	director := newDirector(builder)
	housefromBitBuilder := director.buildHouse()

	fmt.Printf("Bti builders house:: door type %s \n", housefromBitBuilder.doorType)
	fmt.Printf("Bti builders house:: window type %s \n", housefromBitBuilder.windowType)
	fmt.Printf("Bti builders house:: floor %d \n", housefromBitBuilder.floor)
}
