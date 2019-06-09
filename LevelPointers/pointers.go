package main

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
)

// Keyword Identifier Type

//func (r receiver) identifier (parameters) (return){}
// Review interfaces and Method Sets
type tableMeta struct {
	Name       string `json:"name"`
	Connection string `json:"connection"`
	Version    int    `json:"version"`
	Author     string `json:"author"`
	Date       string `json:"date"`
}

func main() {

	x := 78

	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)

	y := &x
	fmt.Println(y)
	fmt.Println(*y)

	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", *y)
	c := circle{5}
	info(&c)

	bd1 := tableMeta{
		Name:       "Usuarios",
		Connection: "SQL=BLABLA",
		Version:    1,
		Author:     "Sharop",
		Date:       "12/12/2019",
	}

	jbd1, err := json.Marshal(bd1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Tipo marshal %T\n", jbd1)
		//fmt.Printf("Sin + %v\n", jbd1)
		//fmt.Printf("con + %+v\n", jbd1)
		fmt.Println(string(jbd1))
	}

	var tmeta tableMeta
	bs := []byte(string(jbd1))
	fmt.Println("Valor de bs, ", bs)

	err = json.Unmarshal(bs, &tmeta)

	if err != nil {
		fmt.Println("Error encontrado, ", err)
	}

	// Unmarshaling we have to take care about wheter is a slice o just an object.
	fmt.Printf("Tipo marshal 2 %T\n", tmeta)
	fmt.Printf("Valor marshal 2 %+v\n", tmeta)

	//fmt.Printf("Sin + %v\n", tmeta)
	//fmt.Printf("con + %+v\n", tmeta)

	people := []Person{
		{"S Bob", 31},
		{"O Jhon", 23},
		{"Hil", 32},
		{"Yes", 34},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}

// Person struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

// ByAge Implements a sort.Interface for []Person
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
