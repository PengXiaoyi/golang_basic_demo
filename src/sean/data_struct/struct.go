package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 25
	return &p
}

type rect struct {
	width, height int
}

// 结构体指针的方法
func (r *rect) area() int {
	return r.width * r.height
}

// 结构体本身的方法，注意结构体和结构体指针本身就是两种不同的类型，但方法也不能声明重名的
// 因为不管是结构体还是结构体指针，都可以调用对方实现的方法
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println(person{"Alice", 20})

	fmt.Println(person{name: "Bob", age: 25})

	fmt.Println(person{name: "Sean"})

	fmt.Println(&person{name: "Ann", age: 26})

	fmt.Println(newPerson("John"))

	s := person{name: "Fred", age: 25}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 26
	fmt.Println(sp.age)

	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim", rp.perim())
}
