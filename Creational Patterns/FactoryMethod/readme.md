
In Go, due to the absence of traditional OOP features like classes and inheritance, we typically use a simplified version of the Factory Method pattern known as the "Simple Factory" pattern.

### Product Interface
```go
type IProduct interface {
	getName() string
	setName(name string)
}
```
We define a Product interface that represents the products common behaviour created by the factory.

### Concrete product implements iProduct
```go
type Product struct {
	name string
}

func (p Product) getName() string {
	return p.name
}

func (p Product) setName(name string) {
	p.name = name
}

```
Created a central product with the implementation of IProduct

### Concrete products
```go
// ConcreteProductA is a type of product.

type ProductA struct {
	Product
}

func newProductA() IProduct {
	return &ProductA{
		Product{
			name: "product a",
		},
	}
}

// ConcreteProductB is a type of product.

type ProductB struct {
	Product
}

func newProductB() IProduct {
	return &ProductB{
		Product{
			name: "product b",
		},
	}

}
```
We create two concrete products, both actually product types. ConcreteProductA and ConcreteProductB both introduce Product with their own characteristics
```go
func (f *Factory) CreateProduct(productType string) IProduct {
	if productType == "product_b" {
		return newProductB()
	} else if productType == "product_a" {
		return newProductA()
	}
	return Product{}
}
```
We define a Factory  responsible for creating products.
We pass product type as argument, each of which creates a specific product type.

### client
```go

func main() {
	fmt.Println("--------------Product A-----------------")
	factory := Factory{}
	productA := factory.CreateProduct("product_a")
	fmt.Println("Here is ", productA.getName())
	fmt.Println("--------------Product B------------------")
	productB := factory.CreateProduct("product_b")
	fmt.Println("Here is ", productB.getName())

}



```
Here is the client


###### ref: https://refactoring.guru/design-patterns/builder/go/example#example-0