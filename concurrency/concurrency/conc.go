package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

type iFixer interface {
	Do() int
}

type plumber struct {
	name       string
	age        int
	speciality []string
}

func (p plumber) Do() int {
	fmt.Println("Ya hice")
	return 0
}

func inf(f iFixer) int {
	return f.Do()
}
func main() {
	fmt.Println("Printing routines")

	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("Arch:\t\t", runtime.GOARCH)
	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("ROUTINES:\t", runtime.NumGoroutine())

	wg.Add(1)

	go foo()
	go bar()
	mbros := plumber{
		name: "Mario Bross",
		age:  45,
		speciality: []string{
			"Drenaje profundo",
			"plantas",
		},
	}

	mbros.Do()
	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("ROUTINES:\t", runtime.NumGoroutine())
	wg.Wait()

	// inf
	fmt.Println("La salida de la interfaz", inf(mbros))

	fmt.Printf("%T\n", wg)
	fmt.Printf("%T\n", &wg)
	fmt.Printf("%T\n", (&wg))

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar: ", i)
	}
}
