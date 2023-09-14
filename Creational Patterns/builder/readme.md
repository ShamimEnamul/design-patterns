## Builder is a creational design pattern, which allows constructing complex objects step by step.
### Product
```go

type House struct {
	windowType string
	doorType   string
	floor      int
}

```
Here is our product. what we want to build and the implementation is colmplex that's why we introduce the builder design pattern


### Builder Interface
```go
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

```
This builder interface make sure the common behavior of the upcoming concrete behavior

```go
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
```

Here, we introduce two concrete builder. They implemented all behavior instructed in IBuilder. Besides, they may have additional behavior but we will not bother with that for now





### Director
```go

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

func (d *Director) buildHouse() House {
d.builder.setDoorType()
d.builder.setNumFloor()
d.builder.setDoorType()
return d.builder.getHouse()
}
```
here, director work as a specific builders director. then the director build houses with the builder


### client
```go
func main() {
	builder := getBuilder("btibuilder") //newBtiBuilder()
	director := newDirector(builder)
	housefromBitBuilder := director.buildHouse()

	fmt.Printf("Bti builders house:: door type %s \n", housefromBitBuilder.doorType)
	fmt.Printf("Bti builders house:: window type %s \n", housefromBitBuilder.windowType)
	fmt.Printf("Bti builders house:: floor %d \n", housefromBitBuilder.floor)
}
```


###### ref: https://refactoring.guru/design-patterns/builder/go/example#example-0