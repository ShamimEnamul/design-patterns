package main

import "fmt"

// Example 1

type LegacyPrinter struct {
}

func (l *LegacyPrinter) print(s string) {

	fmt.Printf("legacy printer -> ", s)
}

type Adapter struct {
	legacyPrinter *LegacyPrinter
}

type ModernPrinter interface {
	printNew(s string)
}

func (a *Adapter) printNew(s string) {
	fmt.Println("modern printer -> ", s)
	a.legacyPrinter.print(s)
}

func main() {
	l := &LegacyPrinter{}
	a := Adapter{l}
	a.printNew("hello")

}
