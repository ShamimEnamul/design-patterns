### Adapter
The Adapter acts as a wrapper between two objects. It catches calls for one object 
and transforms them to format and interface recognizable by the second object.

**Target**: This is the interface that the client code expects to work with. It defines the specific methods or operations that the client code calls.

**Client**: The client is the code that uses the Target interface to interact with objects.

**Adaptee**: This is an existing class or interface that needs to be integrated into the system but has an interface that is incompatible with the Target interface. The Adaptee is the class that the Adapter will wrap.

**Adapter**: This is the class that implements the Target interface while containing an instance of the Adaptee. It acts as an intermediary, translating the calls from the Target interface into calls that the Adaptee can understand.



### Example 1

```go
package main

import "fmt"

// LegacyPrinter is the legacy code that we want to adapt.
type LegacyPrinter struct {
}
// prints using the legacy code.
func (l *LegacyPrinter)print(s string)  {
	
	fmt.Printf("legacy printer -> ", s)
}
// Adapter adapts LegacyPrinter to the ModernPrinter interface.
type Adapter struct {
	legacyPrinter *LegacyPrinter
}
// ModernPrinter is the new interface we want to use.
type ModernPrinter interface {
	printNew(s string)
}
// PrintNew adapts and calls the legacy Print method.
func (a *Adapter) printNew(s string)  {
	fmt.Println("modern printer -> ", s)
	a.legacyPrinter.print(s)
}

func main()  {
	l := &LegacyPrinter{}
	a := Adapter{l}
	a.printNew("hello")

}
```

- LegacyPrinter is the legacy code with the Print method.
- ModernPrinter is the new interface we want to use.
- Adapter adapts LegacyPrinter to the ModernPrinter interface.
- We create an instance of LegacyPrinter and wrap it with the Adapter.
- We use the ModernPrinter interface with the adapted legacy code to print a message.
This allows us to use the new Modern Printer interface to work with legacy printer





### Example 2