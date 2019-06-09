package service

import (
	"errors"
	"log"
)

type Request struct {
	A, B float64
}

type Response struct {
	Result float64
}

// Operations stands for all the remote fuctions definition in the service.
type Operations interface {
	Addition(*Request, *Response) error
	Subtraction(*Request, *Response) error
}

type Calculator int

func (c *Calculator) Addition(rq *Request, rp *Response) error {

	log.Printf("Executing addition with args: %+v", rq)
	rp.Result = rq.A + rq.B
	return nil

}

func (c *Calculator) Subtraction(rq *Request, rp *Response) error {

	log.Printf("Executing subtraction with args: %+v", rq)
	rp.Result = 0
	return errors.New("unsupported operation")

}
