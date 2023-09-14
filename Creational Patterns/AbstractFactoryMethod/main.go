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

func GetFactory(name string) (IFactory, error) {
	if name == "a" {
		return FactoryA{}, nil
	} else if name == "b" {
		return FactoryB{}, nil
	}
	return nil, fmt.Errorf("no factory found")
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
	factoryA, _ := GetFactory("a")
	productA := factoryA.CreateProduct()
	fmt.Println("Here is ", productA.getName())

	fmt.Println("--------------Factory B------------------")
	factoryB, _ := GetFactory("b")
	productB := factoryB.CreateProduct()
	fmt.Println("Here is ", productB.getName())

}
