package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// var lock = &sync.Mutex{}
type Instance struct {
	name string
}

// var single *Instance
//
//	func getInstanc() *Instance {
//		wg.Done()
//		if single == nil {
//			lock.Lock()
//			defer lock.Unlock()
//			if single == nil {
//				single = &Instance{
//					name: "singleton",
//				}
//				fmt.Println("Just created!")
//				return single
//			} else {
//				fmt.Println("created by different go routine!")
//			}
//		} else {
//			fmt.Println("Already created!")
//		}
//		return single
//	}

// second example
var once sync.Once

var single *Instance

func getInstanc() *Instance {
	wg.Done()
	if single == nil {
		once.Do(func() {
			single = &Instance{
				name: "singleton",
			}
			fmt.Println("Just created!")

		})
	} else {
		fmt.Println("Already created!")
	}
	return single
}

// client
func main() {

	wg.Add(20)
	// Example 2
	for i := 0; i < 20; i++ {
		go getInstanc()

	}
	// example 1
	//for i := 0; i < 20; i++ {
	//	go getInstanc()
	//
	//}
	wg.Wait()
	fmt.Scanln()
}
