package main

import (
	"fmt"
	"math/big"
	"strings"
)

const (
	x         = 32
	y float32 = 29
)
const (
	_ = 2000 + iota
	y1
	y2
	y3
)

// func (receiver) identifier(parameters) (returns(s)) {code}
type news struct {
	title  string
	author string
	year   int
}

type actor struct {
	news
	name string
}

//Method
func (a actor) printData() {
	fmt.Println("Meta over, ", a.title, a.name)
}

func main() {
	sharop := []string{
		`                                                 `,
		`███████╗██╗  ██╗ █████╗ ██████╗  ██████╗ ██████╗ `,
		`██╔════╝██║  ██║██╔══██╗██╔══██╗██╔═══██╗██╔══██╗`,
		`███████╗███████║███████║██████╔╝██║   ██║██████╔╝`,
		`╚════██║██╔══██║██╔══██║██╔══██╗██║   ██║██╔═══╝ `,
		`███████║██║  ██║██║  ██║██║  ██║╚██████╔╝██║     `,
		`╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝     `,
		`                                                 `,
		`                                       🍄🍄🍄   `}

	sign := []string{
		`    _      ____  _      ____  ____  ____    `,
		`   / \  /|/  _ \/ \__/|/  _ \/  _ \/  _ \   `,
		`   | |\ ||| / \|| |\/||| / \|| | \|| / \|   `,
		`   | | \||| \_/|| |  ||| |-||| |_/|| |-||   `,
		`   \_/  \|\____/\_/  \|\_/ \|\____/\_/ \|   `,
		`                                            `,
		`    ____  ____  ____  _  _____ _____ ___  _ `,
		`   / ___\/  _ \/   _\/ \/  __//__ __\\  \// `,
		`   |    \| / \||  /  | ||  \    / \   \  /  `,
		`   \___ || \_/||  \_ | ||  /_   | |   / /   `,
		`   \____/\____/\____/\_/\____\  \_/  /_/    `,
		`                               by S    `}

	defer printLogo(sign)
	fmt.Printf(strings.Repeat(".....", 12))
	fmt.Println(factorialBig(9999))
	printLogo(sharop)

	x := []int{3, 4, 2, 3, 2, 2}
	y := []int{300, 99}

	fmt.Println(x)
	fmt.Println(y)

	x = append(x, y...) // The ... param means the rest of the params.
	fmt.Println(x)

	// Deleting from SLICE
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	f := make([]int, 10, 12)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))

	f[9] = 10
	fmt.Println(f)
	f = append(f, 11)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))
	f = append(f, 12)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))
	f = append(f, 13)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))

	ma := []string{"Sharop", "RedHead", "Canelo"}
	fmt.Println(ma)

	mb := []string{"Sergio", "Haro", "Perez"}
	fmt.Println(mb)

	mn := [][]string{ma, mb}
	fmt.Println(mn)

	m := map[string]int{
		"Somebody":       12,
		"Lester Chester": 45,
	}
	fmt.Println(m)

	if v, ok := m["lolo"]; ok {
		fmt.Println("El valor de ", v)
	} else {
		fmt.Println("no existe")
	}

	if v, ok := m["Somebody"]; ok {
		fmt.Println("El valor de ", v)
	}

	act1 := actor{
		news: news{"El ganso cansado,", "Giorgio Postumo", 2019},
		name: "Me canso ganso Men",
	}
	act1.printData()

	// Passing func sum and array of int
	t := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(t)

	internal := call()
	internalb := call()

	fmt.Println(internal())
	fmt.Println(internalb())
	fmt.Println(internal())
	fmt.Println(internalb())

}

//Function
func printLogo(logo []string) {
	for _, s := range logo {
		fmt.Println(s)
	}
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

// function f   accept variadict  params and return int, and a y variadict params
func evenSum(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}

func call() func() int {
	xinterno := 0
	return func() int {
		fmt.Println("El valor de  xinterno, ", xinterno)
		xinterno++
		return xinterno
	}
}

func factorialBig(number int) string {

	result := big.NewFloat(float64(number))

	one := big.NewFloat(float64(1))

	for i := big.NewFloat(float64(number)); i.Cmp(one) > 0; i = i.Sub(i, one) {

		result = result.Mul(result, i)
	}

	return result.String()

}
