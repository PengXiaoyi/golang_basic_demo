package main

import "fmt"

type EndPoint func(req int) (resp int)

type Middleware func(EndPoint) EndPoint

func NewMul2MW() Middleware {
	return func(next EndPoint) EndPoint {
		return func(param int) int {
			res := next(param)
			return res * 2
		}
	}
}

var Add1 EndPoint = func(param int) int {
	return param + 1
}

func main() {
	Mul2Mw := NewMul2MW()

	param := 2

	res := Mul2Mw(Add1)(param)

	fmt.Println(res)
}
