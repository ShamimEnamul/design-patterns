package main

import "fmt"

// Product Interface

type IProduct interface {
	getName() string
	setName(name string)
}

// Concrete product implements iProduct

type Product struct {
	name string
}

func (p Product) getName() string {
	return p.name
}

func (p Product) setName(name string) {
	p.name = name
}

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

// Factory is an interface for creating products.

type IFactory interface {
	CreateProduct() IProduct
}

// Concrete Factory of IFactory type

type FactoryA struct {
}

func (fa FactoryA) CreateProduct() IProduct {
	return newProductA()
}

func newFactoryA() IFactory {
	return FactoryA{}
}

// Concrete Factory of IFactory type

type FactoryB struct {
}

func (fb FactoryB) CreateProduct() IProduct {
	return newProductB()
}
func newFactoryB() IFactory {
	return FactoryB{}
}

func main() {
	fmt.Println("--------------Factory A------------------")
	factoryA := newFactoryA()
	productA := factoryA.CreateProduct()
	fmt.Println("Here is ", productA.getName())

	fmt.Println("--------------Factory B------------------")
	factoryB := newFactoryB()
	productB := factoryB.CreateProduct()
	fmt.Println("Here is ", productB.getName())

}
