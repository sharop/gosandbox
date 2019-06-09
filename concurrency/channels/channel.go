package main

import "fmt"

func main() {
	fmt.Println("Channels")

	c := make(chan int, 2)
	cReceive := make(<-chan int, 2)
	cSend := make(chan<- int, 2)

	// This works, because dont make a deadlock!
	//go func() { c <- 42 }()
	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----------")
	fmt.Printf("c\t\t%T\n", c)
	fmt.Printf("cReceive\t%T\n", cReceive)
	fmt.Printf("cSend\t\t%T\n", cSend)

}

/*
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source
https://blog.golang.org/pipelines
*/
