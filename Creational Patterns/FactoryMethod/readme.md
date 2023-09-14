
In Go, due to the absence of traditional OOP features like classes and inheritance, we typically use a simplified version of the Factory Method pattern known as the "Simple Factory" pattern.

## In this example:

We define a Product interface that represents the products created by the factory.

We create two concrete product types, ConcreteProductA and ConcreteProductB, which implement the Product interface.

We define a Factory  responsible for creating products.

We pass product type as argument, each of which creates a specific product type.

