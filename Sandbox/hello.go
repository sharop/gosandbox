package main

import (
	"fmt"
	"runtime"
)

var x = 42
var y = "James Bond"
var z = true

type fofo int

var pachu = fofo(43)

const (
	a = iota
	b = iota
	c = iota
)

const (
	_  = iota // gives 0
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	n, err := fmt.Println("A que la ...")
	mola(n, err)
	fmt.Println(`Esta es una prueba de escape para 
  con espacions!!!!
  
  "Comillas"`)

	s := fmt.Sprintf("%v, %v, %v\n", x, y, z)

	fmt.Println(s)

	v := 49
	pachu = fofo(v)
	fmt.Println(pachu)
	fmt.Printf("%v, tipo %T\n", pachu, pachu)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fr := "Yulpecuaro apatzingan 🌱"
	bs := []rune(fr)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T, %T, %T\n", a, b, c)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}

func mola(b int, err error) {
	a, _ := fmt.Println("Caracteres: ", b, " Errores: ", err)

	fmt.Println(a)
}
