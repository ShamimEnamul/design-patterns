
### In this example:
We define a Product interface that represents the products created by the factory.

We create two concrete product types, ConcreteProductA and ConcreteProductB, which implement the Product interface.

We define a Factory interface responsible for creating products.

We implement two concrete factories, ConcreteFactoryA and ConcreteFactoryB, each of which creates a specific product type.

In the main function, we demonstrate how to use these factories to create products. factory, 
which creates guns of the desired type based on an incoming argument. The main.go acts as a client. 
Instead of directly interacting with productA or productB, it relies on productA Factory to create instances of various product of productA type