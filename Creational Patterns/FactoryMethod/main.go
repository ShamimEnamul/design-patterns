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

// Factory for creating products.

type Factory struct {
}

func (f *Factory) CreateProduct(productType string) IProduct {
	if productType == "product_b" {
		return newProductB()
	} else if productType == "product_a" {
		return newProductA()
	}
	return Product{}
}

// client
func main() {
	fmt.Println("--------------Product A-----------------")
	factory := Factory{}
	productA := factory.CreateProduct("product_a")
	fmt.Println("Here is ", productA.getName())
	fmt.Println("--------------Product B------------------")
	productB := factory.CreateProduct("product_b")
	fmt.Println("Here is ", productB.getName())

}
