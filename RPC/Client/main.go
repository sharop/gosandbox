package main

import (
	"inspect/RPC/Server/service"
	"log"
	"net/rpc"
)

// Calculator stands for the RPC client implementation.
type Calculator struct {
	Client *rpc.Client
}

func main() {

	// Connecting to the server
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8767")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	c := &Calculator{Client: client}

	result, err := c.addition(5.6, 3.1)
	if err != nil {
		log.Println("Addition error: " + err.Error())
	} else {
		log.Printf("Addition result: %f", result)
	}

	result, err = c.subtraction(7.8, 11)
	if err != nil {
		log.Println("Subtraction error: " + err.Error())
	} else {
		log.Printf("Subtraction result: %f", result)
	}

}

// addition calls the Addition remote method from the calculator service.
func (c *Calculator) addition(a, b float64) (float64, error) {

	args := service.Request{A: a, B: b}
	var response service.Response
	err := c.Client.Call("Calculator.Addition", args, &response)
	return response.Result, err
}

// subtraction calls the Subtraction remote method from the calculator service.
func (c *Calculator) subtraction(a, b float64) (float64, error) {

	args := service.Request{A: a, B: b}
	var response service.Response
	err := c.Client.Call("Calculator.Subtraction", args, &response)
	return response.Result, err
}
