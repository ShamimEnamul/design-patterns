### Component
all leaf and composite should implement the interface

```go
// Component

type Shape interface {
	draw()
}
```
### leaf
all leaf should implement the component
```go
// leaf

type Circle struct {
	name string
}

func (c *Circle) draw() {
	fmt.Printf("circle called: %s \n", c.name)
}

type Rectangle struct {
	name string
}

func (r *Rectangle) draw() {
	fmt.Printf("circle called: %s \n", r.name)
}
```
### composite :
group of shapes also need to implement the component
```go


type GroupOfShapes struct {
	shapes []Shape
	name   string
}

func (g *GroupOfShapes) draw() {
	fmt.Printf("Group of shapes called! \n")
	for _, shape := range g.shapes {
		shape.draw()
	}
}
func (g *GroupOfShapes) add(s Shape) {
	g.shapes = append(g.shapes, s)
}
```
### client
```go
// client
func main() {
	c1 := &Circle{
		name: "football shaped",
	}
	r1 := &Rectangle{
		name: "rectangle shaped",
	}

	g1 := &GroupOfShapes{
		name:   "shapes for 116",
		shapes: []Shape{c1, r1},
	}
	fmt.Println("G1 draw called!")
	g1.draw()

	g2 := GroupOfShapes{
		name:   "shapes for 117",
		shapes: []Shape{g1},
	}
	fmt.Println(g2.name)
	g2.draw()

	g3 := GroupOfShapes{
		name: "shapes for 117",
	}
	g3.add(&g2)

	fmt.Println(g3.name)
	g3.draw()
}
```


### Example 2:
Read from here, https://refactoring.guru/design-patterns/composite/go/example#example-0
