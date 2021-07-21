package main

import (
	"fmt"
)

func main() {
	a := A{sonA: 1}
	b := B{
		sonB: 1,
	}
	c := C{a: a, sonC: 1}
	b.A = a
	fmt.Printf("a:%v ,b:%v ,c:%v", a, b, c)
}

func test1() (err error) {
	a, err := test2()
	fmt.Println(a)
	return
}

func test2() (a int, err error) {
	return 1, &TestError{}
}

type TestError struct {
	Err error // the reason the conversion failed (e.g. ErrRange, ErrSyntax, etc.)
}

func (e *TestError) Error() string {
	return "testerr"
}

type A struct {
	sonA int
}

type B struct {
	A
	sonB int
}

type C struct {
	a    A
	sonC int
}
