package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%v is %v\n", number, "even")
		} else {
			fmt.Printf("%v is %v\n", number, "odd")

		}
	}

}

/*
Here are the directions:

Create a new project folder to house this new project and create a 'main.go' file inside of it.
In the main.go file, define a function called 'main'.  Remember that the 'main' function will be called automatically.
In the main function, create a slice of ints from 0 through 10
Iterate through the slice with a for loop.  For each element, check to see if the number is even or odd.
If the value is even, print out "even".  If it is odd, print out "odd"
Run your code from the terminal by changing into your project directory then running 'go run main.go'.*/
